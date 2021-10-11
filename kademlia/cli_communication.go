package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"strings"
	"sync"
)

func Cli_Channel(state Kademlia, stateMutex *sync.Mutex, cliCh chan string, debug bool) {
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
