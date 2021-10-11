package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"program/kademlia/msg"
	"program/udp"
	"strings"
	"sync"
)

var state Kademlia

func Run(st Kademlia, cliCh chan string) {
	state = st
	/*Docker should always create the node 10.20.0.2
	use that as bootstrapping node to start the network*/
	var stateMutex = &sync.Mutex{}
	if udp.GetOutboundIP().String() != "10.20.0.2" {
		ip := "10.20.0.2:1234"
		stateMutex.Lock()
		public := NewContact(NewSha1KademliaID([]byte(ip)), ip)
		//Add public node to routing table
		state.routingTable.AddContact(public)
		//The perform the join rpc
		joinlookup := NewLookUp(k, "JOIN", []byte(""))
		myid := NewSha1KademliaID([]byte(state.network.Self))
		convID, _ := uuid.NewV4()
		state.convIDMap[*convID] = *joinlookup
		state.network.SendFindContactMessage(&public, *convID, *myid)
		stateMutex.Unlock()
	}
	debug := false
	for {
		select {
		case recv, serverChStatus := <-state.network.RecvRPC:
			if serverChStatus {
				/*Incoming correct messages to the server should
				have the sender added to routing table*/
				conid := NewSha1KademliaID([]byte(recv.Address))
				stateMutex.Lock()
				state.routingTable.AddContact(NewContact(conid, recv.Address))
				stateMutex.Unlock()
				if debug {
					fmt.Println(recv)
				}
				switch recv.RPC {
				case "PING":
					data := msg.MakePong(state.network.Self, recv.ConvID)
					udp.Client(recv.Address, data)
				case "PONG":
					//Idk
				case "FIND_CONTACT":
					//Find the k closest nodes and send them back
					kadID := NewKademliaID(recv.TargetID)
					target := NewContact(kadID, "")
					stateMutex.Lock()
					contacts := state.KClosestNodes(&target)
					var addressList []string
					for _, v := range contacts {
						addressList = append(addressList, v.Address)
					}
					response := msg.MakeFindContactResponse(state.network.Self, addressList,
						recv.TargetID, recv.ConvID, "")
					udp.Client(recv.Address, response)
					stateMutex.Unlock()
				case "FIND_VALUE":
					key := recv.TargetID
					stateMutex.Lock()
					val, ok := state.valueMap[key]
					if ok {
						//If value is found send back the value and an empty contact list
						var addressList []string
						response := msg.MakeFindContactResponse(state.network.Self, addressList,
							recv.TargetID, recv.ConvID, string(val))

						udp.Client(recv.Address, response)
					} else {
						//Otherwise, standard find_contact procedure
						kadID := NewKademliaID(recv.TargetID)
						target := NewContact(kadID, "")
						contacts := state.KClosestNodes(&target)
						var addressList []string
						for _, v := range contacts {
							addressList = append(addressList, v.Address)
						}
						response := msg.MakeFindContactResponse(state.network.Self, addressList,
							recv.TargetID, recv.ConvID, "")
						udp.Client(recv.Address, response)
					}
					stateMutex.Unlock()
				case "FIND_CONTACT_RESPONSE":
					stateMutex.Lock()
					lookup, ok := state.convIDMap[recv.ConvID]
					if !ok {
						break
					}
					//If find_value has found a value we mark the lookup as done and save the value
					if recv.Value != "" && lookup.rpctype == "GET" {
						lookup.foundValue = true
						lookup.value = []byte(recv.Value)
					}

					addrList := recv.Contacts
					contactList := NewResultList(k)
					targetID := NewKademliaID(recv.TargetID)
					//Create a resultlist for the new contacts
					for _, v := range addrList {
						kadID := NewSha1KademliaID([]byte(v))
						contact := NewContact(kadID, v)
						contactList.Insert(contact, *targetID)
					}
					//Merge new contacts into old list and update map
					lookup.klist.Merge(contactList, *targetID)
					//Sort the list so the best nodes are picked
					lookup.klist.Sort(*targetID)
					receivedID := NewSha1KademliaID([]byte(recv.Address))

					lookup.sentmap[receivedID.String()] = true
					state.convIDMap[recv.ConvID] = lookup
					//k new find_nodes need to be sent
					count := 0
					fmt.Println()
					for _, v := range lookup.klist.List {
						if lookup.foundValue { //Value is found, don't send any more requests
							break
						}
						if _, ok := lookup.sentmap[v.ID.String()]; !ok {
							//if nil
							//Send find node
							rpc := msg.MakeFindContact(state.network.Self, targetID.String(), recv.ConvID)
							udp.Client(v.Address, rpc)
							lookup.sentmap[v.ID.String()] = false //No response yet
							go Lookup_timer(stateMutex, v, recv.ConvID, state)
							count++
						}
						if count >= alpha {
							break
						}
					}

					state.convIDMap[recv.ConvID] = lookup //Update map before checking if done

					CheckMsgChainDone(lookup, recv.ConvID)
					stateMutex.Unlock()

				case "STORE":
					stateMutex.Lock()
					state.Store(recv.StoreValue)
					stateMutex.Unlock()
				}
			} else {
				fmt.Println("Channel closed")
			}
		default:
			//idk
		}
		select {
		case cliInst, cliChStatus := <-cliCh:
			if cliChStatus {
				n := strings.Index(cliInst, "|")
				switch cliInst[:n] {
				case "ping":
					reciever := NewContact(NewSha1KademliaID([]byte(cliInst[n+1:])), cliInst[n+1:])
					state.network.SendPingMessage(&reciever)
				case "find closest":
					contacts := state.routingTable.FindClosestContacts(state.routingTable.me.ID, 1)
					for i := range contacts {
						fmt.Printf("%v\n", contacts[i].String())
					}
				case "print buckets":
					for _, bucket := range state.routingTable.buckets {
						for e := bucket.list.Front(); e != nil; e = e.Next() {
							fmt.Printf("%v\n", e.Value)
						}
					}
				case "add contact":
					id := NewSha1KademliaID([]byte(cliInst[n+1:]))
					c := NewContact(id, cliInst[n+1:])
					stateMutex.Lock()
					state.routingTable.AddContact(c)
					stateMutex.Unlock()
				case "put":
					storeVal := cliInst[n+1:]
					stateMutex.Lock()
					storeLookup := NewLookUp(k, "STORE", []byte(storeVal))
					convID, _ := uuid.NewV4()
					storeTarget := NewSha1KademliaID([]byte(storeVal))
					state.convIDMap[*convID] = *storeLookup
					state.LookupContact(storeTarget, *convID, stateMutex)
					stateMutex.Unlock()
				case "get":
					key := cliInst[n+1:]
					stateMutex.Lock()
					getLookup := NewLookUp(k, "GET", nil)
					convID, _ := uuid.NewV4()
					target := NewKademliaID(key)
					state.convIDMap[*convID] = *getLookup
					state.LookupData(target, *convID, stateMutex)
					stateMutex.Unlock()
				case "map":
					stateMutex.Lock()
					fmt.Println(state.valueMap)
					stateMutex.Unlock()
				case "debug":
					debug = !debug
				case "delete":
					reciever := NewContact(NewSha1KademliaID([]byte(cliInst[n+1:])), cliInst[n+1:])
					stateMutex.Lock()
					state.routingTable.deleteContact(reciever)
					stateMutex.Unlock()
				default:
					fmt.Println("Unknown command")
				}
			} else {
				fmt.Println("Channel closed")
			}
		default:
			//idk
		}
	}
}

func CheckMsgChainDone(lookup LookUp, ConvID uuid.UUID) {
	count := 0
	for _, responded := range lookup.sentmap {
		if !responded {
			count++
		}
	}

	if count == 0 {
		//All contacts have responded, we are done
		if lookup.rpctype == "STORE" {
			//Instruct the nodes to store
			for _, v := range lookup.klist.List {
				state.network.SendStoreMessage(lookup.value, v.Address)
			}
		}
		if lookup.rpctype == "JOIN" {
			for _, v := range lookup.klist.List {
				joinid := NewSha1KademliaID([]byte(v.Address))
				state.routingTable.AddContact(NewContact(joinid, v.Address))
			}
		}
		if lookup.rpctype == "GET" {
			if !lookup.foundValue {
				fmt.Println("The value for the given key was not found.")
			} else {
				fmt.Println("The value for the given key is: ", string(lookup.value))
			}
		}
		//Delete the lookup when we are done with the conversation
		delete(state.convIDMap, ConvID)
	}

}
