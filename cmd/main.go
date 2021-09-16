package main

import (
	"bufio"
	"fmt"
	"os"
	k "program/kademlia"
	"program/kademlia/msg"
	"program/udp"
	"strings"
	"time"
)

func main() {
	port := "1234"
	address := udp.GetOutboundIP().String() + ":" + port
	me := k.NewContact(k.NewRandomKademliaID(), address)
	serverCh := make(chan msg.RPC, 50)
	cliCh := make(chan string, 50)
	node := k.NewKademlia(me, serverCh)
	go k.Run(*node, cliCh)
	reader := bufio.NewReader(os.Stdin)
	run := true
	time.Sleep(time.Duration(1000) * time.Millisecond)
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
			cliCh <- "ping|" + address
		default:
			fmt.Printf("Not a command\n")
		}
	}
}
