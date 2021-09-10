package src

const alpha = 3
const k = 20

type Kademlia struct {
	routingTable RoutingTable
	network Network
}

func NewKademlia(me Contact) *Kademlia{
	kademlia := &Kademlia{}
	rt := NewRoutingTable(me)
	kademlia.routingTable = *rt
	network := &Network{}
	kademlia.network = *network
	return kademlia
}

// LookupContact performs the node-lookup recursively
func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
	closestContacts := kademlia.routingTable.FindClosestContacts(target.ID, alpha)
	for i := 0; i < len(closestContacts); i++ {
		kademlia.network.SendFindContactMessage(&closestContacts[i])
	}
}

// KClosestNodes Finds the k closest nodes to the target
func (kademlia *Kademlia) KClosestNodes(target *Contact) []Contact{
	closestContacts := kademlia.routingTable.FindClosestContacts(target.ID, k)
	return closestContacts
}



func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
