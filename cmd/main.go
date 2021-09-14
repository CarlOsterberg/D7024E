package main

import (
	"bufio"
	"fmt"
	"os"
	//json "github.com/CarlOsterberg/D7024E/kademlia/json"
	//kademlia "github.com/CarlOsterberg/D7024E/kademlia"
	//udp "github.com/CarlOsterberg/D7024E/udp"
	kademlia "program/kademlia"
	udp "program/udp"
	"strings"
)

func main() {
	ntwrk := kademlia.Network{}
	go ntwrk.Listen("0.0.0.0", 1234)
	ntwrk.Self = udp.GetOutboundIP().String() + ":1234"
	reader := bufio.NewReader(os.Stdin)
	run := true
	for run {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		switch strings.Replace(text, "\n", "", -1) {
		case "self":
			fmt.Printf("%v\n", udp.GetOutboundIP())
		case "q":
			run = false
		case "ping":
			fmt.Print("Address: ")
			address, _ := reader.ReadString('\n')
			address = strings.Replace(address, "\n", "", -1)
			contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), address)
			ntwrk.SendPingMessage(&contact)
		case "send":
			fmt.Print("Address: ")
			ip, _ := reader.ReadString('\n')
			ip = strings.Replace(ip, "\n", "", -1)
			fmt.Print("Msg: ")
			msg, _ := reader.ReadString('\n')
			msg = strings.Replace(msg, "\n", "", -1)
			udp.Client(ip, msg)
		default:
			fmt.Printf("Not a command\n")
		}
	}
}
