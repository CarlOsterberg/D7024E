package json

import (
	"encoding/json"
	"fmt"
)

//TODO add kademlia id
type RPC struct {
	RPC string
	//uuid *uuid.UUID
	Address string
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
	//u, _ := uuid.NewV4()
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
