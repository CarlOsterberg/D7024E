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
		ConvID:  decodedFormatted.ConvID,
		Address: "12.34.56.78:1234",
	}
	assert.Equal(t, decodedFormatted, correct, "MakePing() error")
}

func TestMakePong(t *testing.T) {
	u, _ := uuid.NewV4()
	formatted := MakePong("12.34.56.78:1234", *u)
	correct := &RPC{
		RPC:     "PONG",
		ConvID:  *u,
		Address: "12.34.56.78:1234",
	}
	decodedFormatted := DecodeRPC(formatted)
	assert.Equal(t, decodedFormatted, correct, "MakePong() error")
}

func TestMakeFindContact(t *testing.T) {
	u, _ := uuid.NewV4()
	formatted := MakeFindContact("12.34.56.78:1234", "00000000000000000000", *u)
	correct := &RPC{
		RPC: "FIND_CONTACT",
		ConvID: *u,
		Address: "12.34.56.78:1234",
		TargetID: "00000000000000000000",
	}
	decodedFormatted := DecodeRPC(formatted)
	assert.Equal(t, decodedFormatted, correct, "MakeFindContactError")
}

func TestMakeFindContactResponse(t *testing.T) {
	u, _ := uuid.NewV4()
	var idList []string
	idList = append(idList, "hello")
	idList = append(idList, "ok")
	formatted := MakeFindContactResponse("12.34.56.78:1234", idList, "00000000000000000000", *u)
	correct := &RPC{
		RPC: "FIND_CONTACT_RESPONSE",
		ConvID: *u,
		Address: "12.34.56.78:1234",
		Contacts: idList,
		TargetID: "00000000000000000000",
	}
	decodedFormatted := DecodeRPC(formatted)
	assert.Equal(t, decodedFormatted, correct, "MakeFindContactError")
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
