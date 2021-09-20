package kademlia

import (
	"crypto/sha1"
	"program/kademlia/msg"
	"strconv"
	"strings"
)

const alpha = 3
const k = 20

type Kademlia struct {
	routingTable RoutingTable
	valueMap     map[string][]byte
	network      Network
	responseMap  map[string]string
}

func NewKademlia(me Contact, ch chan msg.RPC) *Kademlia {
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
	go kademlia.network.Listen("0.0.0.0", port)
	valueMap := make(map[string][]byte)
	kademlia.valueMap = valueMap
	return kademlia
}

// LookupContact performs the node-lookup recursively
func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
	closestContacts := kademlia.routingTable.FindClosestContacts(target.ID, alpha)
	for i := 0; i < len(closestContacts); i++ {
		kademlia.network.SendFindContactMessage(&closestContacts[i])
	}

	//u, err := uuid.NewV4()
}

// KClosestNodes Finds the k closest nodes to the target
func (kademlia *Kademlia) KClosestNodes(target *Contact) []Contact {
	closestContacts := kademlia.routingTable.FindClosestContacts(target.ID, k)
	return closestContacts
}

func (kademlia *Kademlia) LookupData(hash string) []byte {

	//fmt.Println("map:", string(kademlia.valueMap[hash]))
	return kademlia.valueMap[hash]
}

func (kademlia *Kademlia) Store(data []byte) {
	// https://gobyexample.com/sha1-hashes
	// https://gobyexample.com/maps
	key := sha1.New()
	key.Write(data)
	var id = string(key.Sum(nil))
	kademlia.valueMap[id] = data
}
