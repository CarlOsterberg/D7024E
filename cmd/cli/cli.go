package cli

import (
	"bufio"
	"fmt"
	"os"
	"program/udp"
	"strings"
)

func CLI(cliCh chan string) {
	reader := bufio.NewReader(os.Stdin)
	run := true
	//cli loop
	for run {
		text, _ := reader.ReadString('\n')
		switch strings.Replace(text, "\n", "", -1) {
		case "self":
			fmt.Printf("%v\n", udp.GetOutboundIP())
		case "exit":
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
		case "map":
			cliCh <- "map|"
		case "debug":
			cliCh <- "debug|"
		default:
			fmt.Printf("Not a command\n")
		}
	}
}
