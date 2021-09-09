package main

import (
	"bufio"
	"fmt"
	"os"
	//_ "github.com/CarlOsterberg/D7024E/kademlia"
	udp "github.com/CarlOsterberg/D7024E/udp"
	"strings"
)

func main() {
	comms := make(chan string)
	go udp.Server(comms)
	reader := bufio.NewReader(os.Stdin)
	run := true
	for run {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		switch strings.Replace(text, "\n", "", -1) {
		case "send":
			fmt.Print("IP: ")
			ip, _ := reader.ReadString('\n')
			ip = strings.Replace(ip, "\n", "", -1)
			fmt.Print("msg: ")
			msg, _ := reader.ReadString('\n')
			msg = strings.Replace(msg, "\n", "", -1)
			udp.Client(ip, msg)
		case "self":
			fmt.Printf("%v\n", udp.GetOutboundIP())
		case "q":
			run = false
		case "ping":
			fmt.Print("IP: ")
			ip, _ := reader.ReadString('\n')
			ip = strings.Replace(ip, "\n", "", -1)
			self := udp.GetOutboundIP().String()
			udp.Client(ip, string(self)+":1234")
		default:
			fmt.Printf("Not a command\n")
		}
	}
}
