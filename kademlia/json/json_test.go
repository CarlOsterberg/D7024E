package json

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	jsonUdp "program/kademlia/json"
	"testing"
)

func TestMakePing(t *testing.T) {
	formatted := jsonUdp.MakePing("12.34.56.78:1234")
	correct := []byte("{\"RPC\":\"PING\",\"Address\":\"12.34.56.78:1234\"}")
	assert.Equal(t, formatted, correct, "MakePing() error")
}

func TestMakePong(t *testing.T) {
	formatted := jsonUdp.MakePong("12.34.56.78:1234")
	correct := []byte("{\"RPC\":\"PONG\",\"Address\":\"12.34.56.78:1234\"}")
	assert.Equal(t, formatted, correct, "MakePong() error")
}

func TestDecodeRPC(t *testing.T) {
	obj := jsonUdp.RPC{
		Address: "0.0.0.0:1234",
	}
	data, err := json.Marshal(obj)
	if err != nil {
		t.Error("Struct not viable")
	}
	decoded := jsonUdp.DecodeRPC(data)
	assert.True(t, jsonUdp.TestRPC(decoded), "TestDecodeRPC()")
}
