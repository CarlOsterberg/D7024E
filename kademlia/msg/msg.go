package msg

import (
	"encoding/json"
	"fmt"
)

//If a field is added make sure all the test works afterwards,
//should be fixed by adding the empty field of the new field to the test json object
type RPC struct {
	RPC     string
	Address string
	Key     string
	Value   string
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
	ping := &RPC{
		RPC:     "PING",
		Address: address,
	}
	data, err := json.Marshal(ping)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func MakePong(address string) []byte {
	pong := &RPC{
		RPC:     "PONG",
		Address: address,
	}
	data, err := json.Marshal(pong)
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
