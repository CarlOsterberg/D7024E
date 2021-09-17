package udp

import (
	"fmt"
	"net"
)

//https://stackoverflow.com/questions/26028700/write-to-client-udp-socket-in-go

func Client(address string, msg []byte) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	conn.Write(msg)
	conn.Close()
}
