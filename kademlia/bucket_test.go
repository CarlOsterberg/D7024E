package kademlia

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBucket_Delete(t *testing.T) {
	buck := newBucket()
	contact1 := NewContact(NewSha1KademliaID([]byte("hej")), "1233")
	contact2 := NewContact(NewSha1KademliaID([]byte("jeh")), "3312")
	buck.AddContact(contact1)
	buck.AddContact(contact2)
	lenght := buck.Len()
	assert.Equal(t, lenght, 2, "MakePing() error")
	buck.Delete(contact1)

}
