package main

import (
	//kademlia "./kademlia"
	"bufio"
	"fmt"
	"github.com/CarlOsterberg/D7024E/udp"
	"os"
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
			udp.Test()
		case "server":
			go func() {
				udp.Server(comms)
			}()
		case "client":
			udp.Client()
		default:
			fmt.Printf("You wrote: %v", text)
		}
	}
}
