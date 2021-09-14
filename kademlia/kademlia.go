package kademlia

import "crypto/sha1"

type Kademlia struct {
	valueMap  map[string][]byte
}
func NewKademlia() *Kademlia{
	kademlia := &Kademlia{}
	valueMap := make(map[string][]byte)
	kademlia.valueMap = valueMap
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) []byte{

	//fmt.Println("map:", string(kademlia.valueMap[hash]))
	return kademlia.valueMap[hash]
}

func (kademlia *Kademlia) Store(data []byte) {
	// https://gobyexample.com/sha1-hashes
	// https://gobyexample.com/maps
	key :=  sha1.New()
	key.Write(data)
	var id = string(key.Sum(nil))
	kademlia.valueMap[id] = data


}
