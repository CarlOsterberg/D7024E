package kademlia

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"program/kademlia/msg"
	"sync"
	"testing"
)

func TestCheckMsgChainDone(t *testing.T) {

	Lookup := NewLookUp(k, "STORE", []byte("storeVal"))
	serverCh := make(chan msg.RPC, 50)
	id := NewContact(NewSha1KademliaID([]byte("123")), "123")
	convID, _ := uuid.NewV4()
	kadem := NewKademlia(id, serverCh, true)
	CheckMsgChainDone(*kadem, *Lookup, *convID)
	Lookup.rpctype = "JOIN"
	CheckMsgChainDone(*kadem, *Lookup, *convID)
	Lookup.rpctype = "GET"
	CheckMsgChainDone(*kadem, *Lookup, *convID)
	Lookup.foundValue = true
	CheckMsgChainDone(*kadem, *Lookup, *convID)
}
func TestServer_Channel(t *testing.T) {
	serverCh := make(chan msg.RPC, 50)
	id := NewContact(NewSha1KademliaID([]byte("123")), "123")
	kadem := NewKademlia(id, serverCh, true)
	var stateMutex = &sync.Mutex{}
	rpc := msg.MakeStore("123", []byte("hej"))
	rpc1 := msg.DecodeRPC(rpc)
	kadem.network.RecvRPC <- *rpc1
	Server_Channel(*kadem, stateMutex, true)
	hash := NewSha1KademliaID([]byte("hej"))
	assert.Equal(t, []byte("hej"), kadem.valueMap[hash.String()])
}
func TestLookup_timer(t *testing.T) {
	serverCh := make(chan msg.RPC, 50)
	id := NewContact(NewSha1KademliaID([]byte("123")), "123")
	//	Lookup := NewLookUp(k, "STORE", []byte("storeVal"))
	kadem := NewKademlia(id, serverCh, true)
	convID, _ := uuid.NewV4()
	//kadem.convIDMap[*convID] = *Lookup
	var stateMutex = &sync.Mutex{}
	Lookup_timer(stateMutex, id, *convID, *kadem)
}
