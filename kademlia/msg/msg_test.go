package msg

import (
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakePing(t *testing.T) {
	formatted := MakePing("12.34.56.78:1234")
	decodedFormatted := DecodeRPC(formatted)
	correct := &RPC{
		RPC:     "PING",
		MsgID:   decodedFormatted.MsgID,
		Address: "12.34.56.78:1234",
	}
	assert.Equal(t, decodedFormatted, correct, "MakePing() error")
}

func TestMakePong(t *testing.T) {
	u, _ := uuid.NewV4()
	formatted := MakePong("12.34.56.78:1234", *u)
	correct := &RPC{
		RPC:     "PONG",
		MsgID:   *u,
		Address: "12.34.56.78:1234",
	}
	decodedFormatted := DecodeRPC(formatted)
	assert.Equal(t, decodedFormatted, correct, "MakePong() error")
}

func TestDecodeRPC(t *testing.T) {
	obj := RPC{
		Address: "0.0.0.0:1234",
	}
	data, err := json.Marshal(obj)
	if err != nil {
		t.Error("Struct not viable")
	}
	decoded := DecodeRPC(data)
	assert.True(t, TestRPC(decoded), "DecodeRPC() error")
}
