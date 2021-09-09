package udp

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

//https://stackoverflow.com/a/26032240
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("Ty 4 msg"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func Server(comms chan string) {
	buffer := make([]byte, 2048)
	//setting IP to nil makes it listen to all available ips
	addr := net.UDPAddr{
		Port: 1234,
		IP:   nil,
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := server.ReadFromUDP(buffer)
		n := bytes.IndexByte(buffer[:], 0)
		fmt.Printf("Server recv: %v from: %s \n", string(buffer[:n]), remoteaddr)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(server, remoteaddr)
	}
}

//https://stackoverflow.com/a/37382208
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
