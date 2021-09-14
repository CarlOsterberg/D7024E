package kademlia

import (
	"bytes"
	"fmt"
	//"github.com/nu7hatch/gouuid" //u, err := uuid.NewV4()
	"log"
	"net"
	json "program/kademlia/json"
	udp "program/udp"
)

type Network struct {
	Self string
}

func (network *Network) Listen(ip string, port int) {
	buffer := make([]byte, 2048)
	server, err := udp.Server(ip, port)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		_, remoteaddr, err := server.ReadFromUDP(buffer)
		n := bytes.IndexByte(buffer[:], 0)
		data := buffer[:n]
		fmt.Printf("%v\n", string(data))
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go network.handleRPC(server, remoteaddr, data)
	}
}

func (network *Network) handleRPC(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	msg := json.DecodeRPC(data)
	if msg != nil {
		switch msg.RPC {
		case "PING":
			_, err := conn.WriteToUDP(json.MakePong(network.Self), addr)
			if err != nil {
				fmt.Println(err)
			}
		default:
			//todo
		}
	}
}

func (network *Network) SendPingMessage(contact *Contact) {
	msg := json.MakePing(contact.Address)
	udp.Client(contact.Address, string(msg))
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
