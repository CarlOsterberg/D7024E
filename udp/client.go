package udp

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

//https://stackoverflow.com/questions/26028700/write-to-client-udp-socket-in-go

func Client(address string, msg string) {
	buffer := make([]byte, 2048)
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, msg)
	_, err = bufio.NewReader(conn).Read(buffer)
	if err == nil {
		n := bytes.IndexByte(buffer[:], 0)
		fmt.Printf("%s\n", string(buffer[:n]))
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
