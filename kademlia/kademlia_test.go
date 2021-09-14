package kademlia

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestKademlia_Store(t *testing.T) {
	var testkad = NewKademlia()
	key :=  sha1.New()
	key.Write([]byte("nej men va fan"))
	testkad.Store([]byte("nej men va fan"))
	key2 :=  sha1.New()
	key2.Write([]byte("tjentjena"))
	testkad.Store([]byte("tjentjena"))
	key3 :=  sha1.New()
	key3.Write([]byte("hej på dig"))
	testkad.Store([]byte("hej på dig"))
	var id = string(key.Sum(nil))
	var id2 = string(key2.Sum(nil))
	var id3 = string(key3.Sum(nil))
	testkad.LookupData(id2)
	fmt.Println("map:", string(testkad.LookupData(id2)))
	fmt.Println("map:", string(testkad.LookupData(id)))
	fmt.Println("map:", string(testkad.LookupData(id2)))
	fmt.Println("map:", string(testkad.LookupData(id3)))

}
