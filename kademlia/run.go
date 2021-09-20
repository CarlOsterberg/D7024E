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
					udp.Client(recv.Address, msg.MakePong(state.network.Self))
				case "PONG":
					fmt.Println(recv)
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
