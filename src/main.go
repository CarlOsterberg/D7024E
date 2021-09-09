package main

import (
	"bufio"
	"fmt"
	"os"
	//kademlia "./kademlia"
	udp "github.com/CarlOsterberg/D7024E/udp"
	"strings"
)

func main() {
	comms := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		switch strings.Replace(text, "\n", "", -1) {
		case "123":
			Test()
		case "server":
			go func() {
				Server(comms)
			}()
		case "client":
			Client()
		default:
			fmt.Printf("You wrote: %v", text)
		}
	}
}
