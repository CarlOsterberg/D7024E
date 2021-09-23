package msg

import (
	"encoding/json"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
)

//If a field is added make sure all the test works afterwards,
//should be fixed by adding the empty field of the new field to the test json object
type RPC struct {
	RPC      string
	MsgID    uuid.UUID
	Address  string
	TargetID string
	Key      string
	Value    string
	Contacts []string
	ConvID   uuid.UUID
}

func TestRPC(x interface{}) bool {
	switch x.(type) {
	case *RPC:
		return true
	case RPC:
		return true
	default:
		return false
	}
}

func MakePing(address string) []byte {
	u, _ := uuid.NewV4()
	ping := &RPC{
		RPC:     "PING",
		MsgID:   *u,
		Address: address,
	}
	data, err := json.Marshal(ping)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func MakePong(address string, msgID uuid.UUID) []byte {
	pong := &RPC{
		RPC:     "PONG",
		MsgID:   msgID,
		Address: address,
	}
	data, err := json.Marshal(pong)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func MakeFindContact(address string, target string) []byte {
	findContact := &RPC{
		RPC:      "FIND_CONTACT",
		Address:  address,
		TargetID: target,
	}
	data, err := json.Marshal(findContact)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func MakeFindContactResponse(address string, list []string) []byte {
	findContactResponse := &RPC{
		RPC:      "FIND_CONTACT_RESPONSE",
		Address:  address,
		Contacts: list,
	}
	data, err := json.Marshal(findContactResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func DecodeRPC(data []byte) *RPC {
	obj := &RPC{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil
	}
	return obj
}
