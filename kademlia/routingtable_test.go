package kademlia

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestKClosest1(t *testing.T){
	kademlia := NewKademlia(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	kademlia.routingTable.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	tc := NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002")

	contacts := kademlia.KClosestNodes(&tc)

	for i := range contacts{
		fmt.Println(contacts[i].String())
	}



}

func TestKClosest2(t *testing.T){
	kademlia := NewKademlia(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	kademlia.routingTable.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	kademlia.routingTable.AddContact(NewContact(NewKademliaID("FFFFFFFF10000000000000000000000000000000"), "localhost:8001"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111110000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111210000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111310000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111410000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("2111111410000000000000000000000000000000"), "localhost:8002"))

	kademlia.routingTable.AddContact(NewContact(NewKademliaID("FFFFFFFF20000000000000000000000000000000"), "localhost:8001"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111120000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111220000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111320000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111420000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("2111111420000000000000000000000000000000"), "localhost:8002"))

	kademlia.routingTable.AddContact(NewContact(NewKademliaID("FFFFFFFF30000000000000000000000000000000"), "localhost:8001"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111130000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111230000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111330000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("1111111430000000000000000000000000000000"), "localhost:8002"))
	kademlia.routingTable.AddContact(NewContact(NewKademliaID("2111111430000000000000000000000000000000"), "localhost:8002"))


	tc := NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002")

	contacts := kademlia.KClosestNodes(&tc)

	for i := range contacts{
		fmt.Println(contacts[i].String())
	}



}