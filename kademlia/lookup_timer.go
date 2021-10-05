package kademlia

import (
	uuid "github.com/nu7hatch/gouuid"
	"sync"
	"time"
)

func Lookup_timer(convIdMap map[uuid.UUID]LookUp, stateMutex *sync.Mutex, kademliaID string, convID uuid.UUID) {
	//timeout
	time.Sleep(10 * time.Second)
	stateMutex.Lock()
	lookup, ok := convIdMap[convID]
	//The conversion chain has ended and this thread should terminate
	if !ok {
		stateMutex.Unlock()
		return
	}
	//If the node hasn't responded yet delete the entry from the sentmap
	if lookup.sentmap[kademliaID] == false {
		delete(lookup.sentmap, kademliaID)
		CheckMsgChainDone(convIdMap[convID], convID)
	}
	stateMutex.Unlock()
}
