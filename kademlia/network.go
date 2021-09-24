package kademlia

import (
	"bytes"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"program/kademlia/msg"
	"program/udp"
)

type Network struct {
	Self    string
	RecvRPC chan msg.RPC
}

func (network *Network) Listen(ip string, port int) {
	server, err := udp.Server(ip, port)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Server running")
	for {
		buffer := make([]byte, 2048)
		_, _, err := server.ReadFromUDP(buffer)
		n := bytes.IndexByte(buffer[:], 0)
		data := buffer[:n]
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go network.decode(data)
	}
}

func (network *Network) decode(data []byte) {
	decodedMsg := msg.DecodeRPC(data)
	if decodedMsg != nil {
		network.RecvRPC <- *decodedMsg
	} else {
		fmt.Println("Msg not valid")
	}
}

func (network *Network) SendPingMessage(contact *Contact) {
	message := msg.MakePing(network.Self)
	udp.Client(contact.Address, message)
}

func (network *Network) SendFindContactMessage(contact *Contact, convID uuid.UUID, target KademliaID) {
	targetID := target.String()
	msg := msg.MakeFindContact(network.Self, targetID, convID)
	udp.Client(contact.Address, msg)

}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte, addres string) {
	storeMessage := msg.MakeStore(data)
	udp.Client(addres, storeMessage)
}
