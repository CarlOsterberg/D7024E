package kademlia

import (
	"fmt"
	"program/kademlia/msg"
	"program/udp"
	"strings"
)

func Run(state Kademlia, cliCh chan string) {
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
				case "STORE":
					fmt.Println(recv)
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
				reciever := NewContact(NewRandomKademliaID(), cliInst[n+1:])
				state.network.SendPingMessage(&reciever)
			} else {
				fmt.Println("Channel closed")
			}
		default:
			//idk
		}
	}
}
