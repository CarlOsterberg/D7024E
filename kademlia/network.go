package kademlia

import (
	"bytes"
	"fmt"
	"log"
	"program/kademlia/msg"
	"program/udp"
)

type Network struct {
	Self    string
	RecvRPC chan msg.RPC
}

func (network *Network) Listen(ip string, port int) {
	buffer := make([]byte, 2048)
	server, err := udp.Server(ip, port)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Server running")
	for {
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
	message := msg.MakePing(contact.Address)
	udp.Client(contact.Address, message)
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	//targetID := contact.ID.String()
	//msg := msg.MakeFindContact(contact.Address, targetID)
	//udp.Client(contact.Address, msg)

}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
