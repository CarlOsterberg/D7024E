package kademlia

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, len(rl.List), 5)

	for _, v := range rl.List {
		fmt.Println(v.ID)
	}

}

func TestInsert2(t *testing.T) {

	rl := NewResultList(5)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")

	rl.Insert(NewContact(id1, "localhost"), *target)
	rl.Insert(NewContact(id2, "localhost"), *target)
	rl.Insert(NewContact(id1, "localhost"), *target)

	assert.Equal(t, 2, len(rl.List))

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

	assert.Equal(t, len(rl1.List), 3)

	for _, v := range rl1.List {
		fmt.Println(v.ID)
	}

}

func TestDelete(t *testing.T) {
	rl1 := NewResultList(3)
	rl2 := NewResultList(3)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")

	rl1.Insert(NewContact(id1, "localhost"), *target)
	rl1.Insert(NewContact(id2, "localhost"), *target)
	rl1.Insert(NewContact(id3, "localhost"), *target)

	rl2.Insert(NewContact(id1, "localhost"), *target)
	rl2.Insert(NewContact(id3, "localhost"), *target)

	rl1.Delete(NewContact(id2, "localhost"))

	for i, _ := range rl1.List {
		assert.Equal(t, rl1.List[i], rl2.List[i])
	}
}

func TestSort(t *testing.T) {

	rl := NewResultList(20)

	target := NewKademliaID("FFFFFFFF00000000000000000000000000000001")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id7 := NewKademliaID("FFFFFFFF00000000000000000000000000000007")
	id8 := NewKademliaID("FFFFFFFF00000000000000000000000000000008")
	id9 := NewKademliaID("FFFFFFFF00000000000000000000000000000009")
	id10 := NewKademliaID("FFFFFFFF00000000000000000000000000000010")
	id11 := NewKademliaID("FFFFFFFF00000000000000000000000000000011")
	id12 := NewKademliaID("FFFFFFFF00000000000000000000000000000012")
	id13 := NewKademliaID("FFFFFFFF00000000000000000000000000000013")
	id14 := NewKademliaID("FFFFFFFF00000000000000000000000000000014")
	id15 := NewKademliaID("FFFFFFFF00000000000000000000000000000015")
	id16 := NewKademliaID("FFFFFFFF00000000000000000000000000000016")
	id17 := NewKademliaID("FFFFFFFF00000000000000000000000000000017")
	id18 := NewKademliaID("FFFFFFFF00000000000000000000000000000018")
	id19 := NewKademliaID("FFFFFFFF00000000000000000000000000000019")
	id20 := NewKademliaID("FFFFFFFF00000000000000000000000000000020")
	id21 := NewKademliaID("FFFFFFFF00000000000000000000000000000021")

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

	rl.Sort(*target)

	assert.Equal(t, len(rl.List), 20)

	var distance = NewKademliaID("0000000000000000000000000000000000000000")

	for _, v := range rl.List {
		assert.False(t, v.ID.CalcDistance(target).Less(distance))
		distance = v.ID.CalcDistance(target)
	}

}
func TestInsert3(t *testing.T) {

	rl := NewResultList(5)

	target := NewKademliaID("FFFFFFFFABCABCFF123445676789ACDF123ABCFA")

	id1 := NewKademliaID("FFFFFFFF00000000000000000000000000000001")
	id2 := NewKademliaID("FFFFFFFF00000000000000000000000000000002")
	id3 := NewKademliaID("FFFFFFFF00000000000000000000000000000003")
	id4 := NewKademliaID("FFFFFFFF00000000000000000000000000000004")
	id5 := NewKademliaID("FFFFFFFF00000000000000000000000000000005")
	id6 := NewKademliaID("FFFFFFFF00000000000000000000000000000006")
	id7 := NewKademliaID("FFFFFFFF00000000000001230000000000000006")
	id8 := NewKademliaID("FFFFFFFF00004560000000000000000000000006")
	id9 := NewKademliaID("FFFFFFFF00000000235500000000000000000006")
	id10 := NewKademliaID("FFFFFFFF00000000234500000000000000000006")
	id11 := NewKademliaID("FFFFFFFF00000000000000678000000000000006")
	id12 := NewKademliaID("FFFFFFFF00000000000000000001230000000006")
	id13 := NewKademliaID("FFFFFFFF00000000000043300000000000000006")
	id14 := NewKademliaID("FFFFFFFF00000000000000000054000000000006")
	id15 := NewKademliaID("FFFFFFFF00000000001200000000000000000006")
	id16 := NewKademliaID("FFFFFFFF00000000000000000430000000000006")
	id17 := NewKademliaID("FFFFFFFF00000012300000000000000000000006")
	id18 := NewKademliaID("FFFFFFFF00000000000000004300000000000006")
	id19 := NewKademliaID("FFFFFFFF00000000003420000000000000000006")
	id20 := NewKademliaID("FFFFFFFF00000000000000012300000000000006")
	id21 := NewKademliaID("FFFFFFFF00000000034200000000000000000006")

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

	assert.Equal(t, len(rl.List), 5)

	for _, v := range rl.List {
		fmt.Println(v.ID)
	}

}
