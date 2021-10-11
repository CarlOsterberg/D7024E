package kademlia

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"program/kademlia/msg"
	"strconv"
	"strings"
	"sync"
)

const alpha = 3
const k = 20

type Kademlia struct {
	routingTable RoutingTable
	valueMap     map[string][]byte
	network      Network
	convIDMap    map[uuid.UUID]LookUp
}

func NewKademlia(me Contact, ch chan msg.RPC, test bool) *Kademlia {
	kademlia := &Kademlia{}
	rt := NewRoutingTable(me)
	kademlia.routingTable = *rt
	network := &Network{
		Self:    me.Address,
		RecvRPC: ch,
	}
	kademlia.network = *network
	n := strings.Index(me.Address, ":")
	port, _ := strconv.Atoi(me.Address[n+1:])
	if !test {
		go kademlia.network.Listen("0.0.0.0", port)
	}
	valueMap := make(map[string][]byte)
	kademlia.valueMap = valueMap
	sentid := make(map[uuid.UUID]LookUp)
	kademlia.convIDMap = sentid
	return kademlia
}

// LookupContact performs the node-lookup recursively
func (kademlia *Kademlia) LookupContact(target *KademliaID, convID uuid.UUID, stateMutex *sync.Mutex) {
	// TODO
	closestContacts := kademlia.routingTable.FindClosestContacts(target, alpha)
	for i := 0; i < len(closestContacts); i++ {
		kademlia.network.SendFindContactMessage(&closestContacts[i], convID, *target)
		kademlia.convIDMap[convID].sentmap[closestContacts[i].ID.String()] = false
		fmt.Printf("Started for %v\n", *target)
		go Lookup_timer(stateMutex, closestContacts[i], convID, state)
	}
}

// KClosestNodes Finds the k closest nodes to the target
func (kademlia *Kademlia) KClosestNodes(target *Contact) []Contact {
	closestContacts := kademlia.routingTable.FindClosestContacts(target.ID, k)
	return closestContacts
}

func (kademlia *Kademlia) LookupData(target *KademliaID, convID uuid.UUID, stateMutex *sync.Mutex) {
	closestContacts := kademlia.routingTable.FindClosestContacts(target, alpha)
	for i := 0; i < len(closestContacts); i++ {
		kademlia.network.SendFindDataMessage(&closestContacts[i], convID, *target)
		kademlia.convIDMap[convID].sentmap[closestContacts[i].ID.String()] = false
		fmt.Printf("Started for %v\n", target)
		go Lookup_timer(stateMutex, closestContacts[i], convID, state)
	}
}

func (kademlia *Kademlia) Store(data []byte) {
	// https://gobyexample.com/sha1-hashes
	// https://gobyexample.com/maps
	id := NewSha1KademliaID(data)

	kademlia.valueMap[id.String()] = data
}
