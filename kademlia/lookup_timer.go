package kademlia

import (
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
		return
	}
	//If the node hasn't responded yet delete the entry from the sentmap
	if lookup.sentmap[contact.ID.String()] == false {
		delete(lookup.sentmap, contact.ID.String())
		lookup.klist.Delete(contact)
		state.convIDMap[convID] = lookup
		state.routingTable.deleteContact(contact)
		CheckMsgChainDone(convMap[convID], convID)
		state.network.SendPingMessage(&contact)
	}
	stateMutex.Unlock()
}
