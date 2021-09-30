package main

import (
	"bufio"
	"fmt"
	"os"
	k "program/kademlia"
	"program/kademlia/msg"
	"program/udp"
	"strings"
)

func main() {
	//startup procedure
	port := "1234"
	address := udp.GetOutboundIP().String() + ":" + port
	//create contact for self
	me := k.NewContact(k.NewSha1KademliaID([]byte(address)), address)
	//channel for server -> node_state communication
	serverCh := make(chan msg.RPC, 50)
	//channel for cli -> node_state communication
	cliCh := make(chan string, 50)
	//create the kademlia network node state
	node := k.NewKademlia(me, serverCh)
	//start the node state thread
	go k.Run(*node, cliCh)
	//start the cli reader
	reader := bufio.NewReader(os.Stdin)
	run := true
	//cli loop
	for run {
		text, _ := reader.ReadString('\n')
		switch strings.Replace(text, "\n", "", -1) {
		case "self":
			fmt.Printf("%v\n", udp.GetOutboundIP())
		case "terminate":
			run = false
		case "ping":
			fmt.Print("Address: ")
			address, _ := reader.ReadString('\n')
			address = strings.Replace(address, "\n", "", -1)
			cliCh <- "ping|" + address
		case "find closest":
			cliCh <- "find closest|"
		case "print buckets":
			cliCh <- "print buckets|"
		case "add contact":
			fmt.Print("Address: ")
			address, _ := reader.ReadString('\n')
			address = strings.Replace(address, "\n", "", -1)
			cliCh <- "add contact|" + address
		case "put":
			fmt.Print("Enter text to store: ")
			data, _ := reader.ReadString('\n')
			data = strings.Replace(data, "\n", "", -1)
			if len(data) < 256 {
				cliCh <- "put|" + data
			} else {
				fmt.Println("Text was to long to be stored (255 characters max)")
			}
		case "get":
			fmt.Print("Enter key: ")
			data, _ := reader.ReadString('\n')
			data = strings.Replace(data, "\n", "", -1)
			if len(data) != 20 {
				fmt.Println("Not a valid key")
			}else{
				cliCh <- "get|" + data
			}
		case "map":
			cliCh <- "map|"
		default:
			fmt.Printf("Not a command\n")
		}
	}
}
