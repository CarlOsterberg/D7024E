package udp

import (
	"fmt"
	"net"
)

//https://stackoverflow.com/questions/26028700/write-to-client-udp-socket-in-go

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) (result string) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
		return "Fail"
	} else {
		return "Success"
	}
}

func Server(comms chan string) {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go func() {
			fmt.Printf("Did work? %v", sendResponse(ser, remoteaddr))
		}()
	}
}
