package kademlia

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewLookUp(t *testing.T) {
	look := NewLookUp(20, "STORE", []byte("hejhej"))
	assert.Equal(t, look.rpctype, "STORE", "lookuprpc error")
	assert.Equal(t, look.value, []byte("hejhej"), "lookuprpc error")
}
