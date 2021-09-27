package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"program/kademlia/msg"
	"program/udp"
	"strings"
)

func Run(state Kademlia, cliCh chan string) {
	if udp.GetOutboundIP().String() != "10.20.0.2" {
		ip := "10.20.0.2:1234"
		public := NewContact(NewSha1KademliaID([]byte(ip)), ip)
		state.routingTable.AddContact(public)
		joinlookup := NewLookUp(k, "JOIN", []byte(""))
		myid := NewSha1KademliaID([]byte(state.network.Self))
		convID, _ := uuid.NewV4()
		state.convIDMap[*convID] = *joinlookup
		state.network.SendFindContactMessage(&public, *convID, *myid)

	}
	for {
		select {
		case recv, serverChStatus := <-state.network.RecvRPC:
			if serverChStatus {
				conid := NewSha1KademliaID([]byte(recv.Address))
				state.routingTable.AddContact(NewContact(conid, recv.Address))
				switch recv.RPC {
				case "PING":
					fmt.Println("inkommande ping")
					fmt.Println(recv)
					udp.Client(recv.Address, msg.MakePong(state.network.Self, recv.ConvID))
				case "PONG":
					fmt.Println("inkommande pong")
					fmt.Println(recv)
				case "FIND_CONTACT":
					fmt.Println("INNE I FIND CONNACT")
					//Find the k closest nodes and send them back
					kadID := NewKademliaID(recv.TargetID)
					target := NewContact(kadID, "")
					contacts := state.KClosestNodes(&target)
					var addressList []string
					for _, v := range contacts {
						addressList = append(addressList, v.Address)
					}
					response := msg.MakeFindContactResponse(state.network.Self, addressList, recv.TargetID, recv.ConvID)
					udp.Client(recv.Address, response)
				case "FIND_CONTACT_RESPONSE":
					fmt.Println("INNE I FIND CONNACT RESPONSE")
					lookup, ok := state.convIDMap[recv.ConvID]
					if !ok {
						panic("ERROR!!!!!!!!!!!! LOOKUP has been DELETED, This line should not be reached")
					}
					addrList := recv.Contacts
					contactList := NewResultList(k)
					targetID := NewKademliaID(recv.TargetID)
					for _, v := range addrList {
						kadID := NewSha1KademliaID([]byte(v))
						contact := NewContact(kadID, v)
						contactList.Insert(contact, *targetID)
					}
					//Merge new contacts into old list and update map
					lookup.klist.Merge(contactList, *targetID)

					//	key := sha1.New()
					//	key.Write([]byte(recv.Address))
					//	receivedID := string(key.Sum(nil))
					//	lookup.sentmap[receivedID] = true
					receivedID := NewSha1KademliaID([]byte(recv.Address))
					lookup.sentmap[receivedID.String()] = true

					state.convIDMap[recv.ConvID] = lookup
					//k new find_nodes need to be sent
					count := 0
					//	fmt.Println("the sent map")
					//fmt.Println(lookup.sentmap)
					for _, v := range lookup.klist.List {
						if _, ok := lookup.sentmap[v.ID.String()]; !ok {
							//if nil
							//Send find node
							rpc := msg.MakeFindContact(state.network.Self, targetID.String(), recv.ConvID)
							udp.Client(v.Address, rpc)
							lookup.sentmap[v.ID.String()] = false //No response yet
							count++
						}
						if count >= alpha {
							break
						}
					}

					state.convIDMap[recv.ConvID] = lookup //Update map before checking if done

					count = 0
					//fmt.Println(lookup.sentmap)
					/*for _, v := range lookup.klist.List {

						if ok, v := lookup.sentmap[v.ID.String()]; ok && v {
							continue
						} else {
							count++
						}
					}*/

					for _, val := range lookup.sentmap {
						if !val {
							count++
						}
					}

					if count == 0 {
						fmt.Println("INNE I NODELOOKUPEND")
						//All contacts have responded, we are done
						if lookup.rpctype == "STORE" {
							fmt.Println("INNE I NODELOOKUPEND Store")
							//Instruct the nodes to store
							for _, v := range lookup.klist.List {
								state.network.SendStoreMessage(lookup.value, v.Address)
							}
						}
						if lookup.rpctype == "JOIN" {
							fmt.Println("INNE I NODELOOKUPEND JOIN")
							fmt.Println(lookup.klist.List)
							for _, v := range lookup.klist.List {
								joinid := NewSha1KademliaID([]byte(v.Address))
								state.routingTable.AddContact(NewContact(joinid, v.Address))
							}
						}

						//Delete the lookup when we are done with the conversation
						delete(state.convIDMap, recv.ConvID)

					}

				case "STORE":
					//fmt.Println(recv)
					state.Store(recv.StoreValue)
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
						fmt.Println(contacts[i].String())
					}
				case "print buckets":
					for _, bucket := range state.routingTable.buckets {
						for e := bucket.list.Front(); e != nil; e = e.Next() {
							fmt.Println(e.Value)
						}
					}
				case "add contact":
					id := NewSha1KademliaID([]byte(cliInst[n+1:]))
					c := NewContact(id, cliInst[n+1:])
					state.routingTable.AddContact(c)
				case "put":
					storeVal := cliInst[n+1:]
					fmt.Print("run.go received: ")
					fmt.Println(storeVal)
					storeLookup := NewLookUp(k, "STORE", []byte(storeVal))
					convID, _ := uuid.NewV4()
					storeTarget := NewSha1KademliaID([]byte(storeVal))
					state.convIDMap[*convID] = *storeLookup
					state.LookupContact(storeTarget, *convID)
				case "map":
					fmt.Println(state.valueMap)
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
