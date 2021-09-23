package kademlia

import (
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
					//lookup := state.convIDMap[recv.ConvID]
					//addrList := recv.Contacts
					//lookup.klist.Merge()

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
