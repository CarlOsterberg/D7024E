package udp

import (
	"log"
	"net"
)

//https://stackoverflow.com/a/26032240
func Server(ip string, port int) (*net.UDPConn, error) {
	//setting IP to nil makes it listen to all available ips
	addr := net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, err
	} else {
		return server, nil
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
