package kademlia

import (
	"fmt"
	"testing"
)

func TestInsert1(t *testing.T){

	rl := NewResultList(5)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")


	rl.Insert(*id1, *target)
	rl.Insert(*id2, *target)
	rl.Insert(*id3, *target)
	rl.Insert(*id4, *target)
	rl.Insert(*id5, *target)
	rl.Insert(*id6, *target)

	for _, v := range rl.List{
		fmt.Println(v)
	}



}

func TestMerge(t *testing.T){
	rl1 := NewResultList(3)
	rl2 := NewResultList( 3)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")

	rl1.Insert(*id1, *target)
	rl1.Insert(*id2, *target)
	rl1.Insert(*id3, *target)
	rl2.Insert(*id4, *target)
	rl1.Insert(*id5, *target)
	rl1.Insert(*id6, *target)

	rl1.Merge(rl2, *target)

	for _, v := range rl1.List{
		fmt.Println(v)
	}

}

