package kademlia

import (
	"fmt"
	"testing"
)

func TestInsert1(t *testing.T) {

	rl := NewResultList(5)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id7 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id8 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id9 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id10 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id11 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id12 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id13 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id14 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id15 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id16 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id17 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id18 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id19 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id20 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id21 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")

	rl.Insert(NewContact(id1, "localhost"), *target)
	rl.Insert(NewContact(id2, "localhost"), *target)
	rl.Insert(NewContact(id3, "localhost"), *target)
	rl.Insert(NewContact(id4, "localhost"), *target)
	rl.Insert(NewContact(id5, "localhost"), *target)
	rl.Insert(NewContact(id6, "localhost"), *target)
	rl.Insert(NewContact(id7, "localhost"), *target)
	rl.Insert(NewContact(id8, "localhost"), *target)
	rl.Insert(NewContact(id9, "localhost"), *target)
	rl.Insert(NewContact(id10, "localhost"), *target)
	rl.Insert(NewContact(id11, "localhost"), *target)
	rl.Insert(NewContact(id12, "localhost"), *target)
	rl.Insert(NewContact(id13, "localhost"), *target)
	rl.Insert(NewContact(id14, "localhost"), *target)
	rl.Insert(NewContact(id15, "localhost"), *target)
	rl.Insert(NewContact(id16, "localhost"), *target)
	rl.Insert(NewContact(id17, "localhost"), *target)
	rl.Insert(NewContact(id18, "localhost"), *target)
	rl.Insert(NewContact(id19, "localhost"), *target)
	rl.Insert(NewContact(id20, "localhost"), *target)
	rl.Insert(NewContact(id21, "localhost"), *target)
	for _, v := range rl.List {
		fmt.Println(v.ID)
	}

}

func TestMerge(t *testing.T) {
	rl1 := NewResultList(3)
	rl2 := NewResultList(3)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")

	rl1.Insert(NewContact(id1, "localhost"), *target)
	rl1.Insert(NewContact(id2, "localhost"), *target)
	rl1.Insert(NewContact(id3, "localhost"), *target)
	rl2.Insert(NewContact(id4, "localhost"), *target)
	rl2.Insert(NewContact(id5, "localhost"), *target)
	rl2.Insert(NewContact(id6, "localhost"), *target)

	rl1.Merge(rl2, *target)

	for _, v := range rl1.List {
		fmt.Println(v.ID)
	}

}
