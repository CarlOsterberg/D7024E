package kademlia

import (
	"fmt"
	"program/kademlia/msg"
	"program/udp"
	"sync"
)

func Server_Channel(state Kademlia, stateMutex *sync.Mutex, debug bool) {
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

				CheckMsgChainDone(state, lookup, recv.ConvID)
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
}
