package kademlia

import (
	"github.com/stretchr/testify/assert"
	"program/kademlia/msg"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestKademlia_Store(t *testing.T) {
	serverCh := make(chan msg.RPC, 50)
	var testkad = NewKademlia(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"), serverCh, true)
	key1 := NewSha1KademliaID([]byte("nej men va fan"))
	key2 := NewSha1KademliaID([]byte("tjentjena"))
	key3 := NewSha1KademliaID([]byte("hej på dig"))
	value1 := "nej men va fan"
	value2 := "tjentjena"
	value3 := "hej på dig"
	testkad.Store([]byte(value1))
	testkad.Store([]byte(value2))
	testkad.Store([]byte(value3))
	assert.Equal(t, value1, string(testkad.valueMap[key1.String()]), "value1 store error")
	assert.Equal(t, value2, string(testkad.valueMap[key2.String()]), "value2 store error")
	assert.Equal(t, value3, string(testkad.valueMap[key3.String()]), "value3 store error")
}
