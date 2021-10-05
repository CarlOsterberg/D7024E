package main

import (
	"program/cmd/cli"
	k "program/kademlia"
	"program/kademlia/msg"
	"program/udp"
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
	node := k.NewKademlia(me, serverCh, false)
	//start the node state thread
	go k.Run(*node, cliCh)
	cli.CLI(cliCh)
}
