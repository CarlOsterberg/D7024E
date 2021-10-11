package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"sync"
	"time"
)

func Lookup_timer(stateMutex *sync.Mutex, contact Contact, convID uuid.UUID, state Kademlia) {
	//timeout
	time.Sleep(10 * time.Second)
	stateMutex.Lock()
	convMap := state.convIDMap
	lookup, ok := convMap[convID]
	//The conversion chain has ended and this thread should terminate
	if !ok {
		stateMutex.Unlock()
		fmt.Printf("Abort %v\n", contact.ID.String())
		return
	}
	fmt.Println("Pre if, lookup_timer.go")
	fmt.Println(contact.ID.String())
	for i, v := range lookup.sentmap {
		fmt.Printf("%v status %v\n", i, v)
	}
	//If the node hasn't responded yet delete the entry from the sentmap
	if lookup.sentmap[contact.ID.String()] == false {
		delete(lookup.sentmap, contact.ID.String())
		state.routingTable.deleteContact(contact)
		fmt.Printf("attempt to delete %v\n", contact)
		CheckMsgChainDone(convMap[convID], convID)
		state.network.SendPingMessage(&contact)
	}
	stateMutex.Unlock()
}
