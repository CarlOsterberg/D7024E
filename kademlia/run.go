package kademlia

import (
	"crypto/sha1"
	"fmt"
	"program/kademlia/msg"
	"program/udp"
	"strings"
)

func Run(state Kademlia, cliCh chan string) {
	if udp.GetOutboundIP().String() != "172.20.0.2" {
		ip := "172.20.0.2:1234"
		public := NewContact(NewSha1KademliaID([]byte(ip)), ip)
		state.routingTable.AddContact(public)
	}
	for {
		select {
		case recv, serverChStatus := <-state.network.RecvRPC:
			if serverChStatus {
				switch recv.RPC {
				case "PING":
					fmt.Println(recv)
					udp.Client(recv.Address, msg.MakePong(state.network.Self, recv.MsgID))
				case "PONG":
					fmt.Println(recv)
				case "FIND_CONTACT":
					//Find the k closest nodes and send them back
					kadID := NewKademliaID(recv.TargetID)
					target := NewContact(kadID, "")
					contacts := state.KClosestNodes(&target)
					var addressList []string
					for _, v := range contacts {
						addressList = append(addressList, v.Address)
					}
					response := msg.MakeFindContactResponse(state.network.Self, addressList)
					udp.Client(recv.Address, response)
				case "FIND_CONTACT_RESPONSE":
					lookup := state.convIDMap[recv.ConvID]
					addrList := recv.Contacts
					contactList := NewResultList(k)
					targetID := NewKademliaID(recv.TargetID)
					for _, v := range addrList{
						//Hash the addresses and insert contacts into a list
						key := sha1.New()
						key.Write([]byte(v))
						id := string(key.Sum(nil))
						kadID := NewKademliaID(id)
						contact := NewContact(kadID, v)
						contactList.Insert(contact, *targetID)
					}
					//Merge new contacts into old list and update map
					lookup.klist.Merge(contactList, *targetID)

					key := sha1.New()
					key.Write([]byte(recv.Address))
					receivedID := string(key.Sum(nil))
					lookup.sentmap[receivedID] = true

					state.convIDMap[recv.ConvID] = lookup
					//k new find_nodes need to be sent
					count := 0
					for _, v := range lookup.klist.List {
						if _, ok := lookup.sentmap[v.ID.String()]; !ok{
							//if nil
							//Send find node
							rpc := msg.MakeFindContact(state.network.Self, targetID.String())
							udp.Client(v.Address, rpc)
							lookup.sentmap[v.ID.String()] = false //No response yet
							count++
						}
						if count >= alpha{
							break
						}
					}

					state.convIDMap[recv.ConvID] = lookup //Update map before checking if done

					done := false
					count = 0
					for _, v := range lookup.klist.List {
						if ok, v := lookup.sentmap[v.ID.String()]; ok && v{
							continue
						} else {
							count++
						}
					}

					if count == 0 {
						done = true
					}

					if done{
						//All contacts have responded, we are done
						if lookup.rpctype == "STORE"{
							//Instruct the nodes to store
						}


						//TODO delete
					}

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
					reciever := NewContact(NewRandomKademliaID(), cliInst[n+1:])
					state.network.SendPingMessage(&reciever)
				case "find closest":
					contacts := state.routingTable.FindClosestContacts(state.routingTable.me.ID, 1)
					for i := range contacts {
						fmt.Println(contacts[i].String())
					}
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
