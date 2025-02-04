package messageserver

import (
	"fmt"
	"net"
)

var UdpClient *net.UDPConn

func ConnectToUdp() {
	serverAddr := "18.185.47.252:5000"

	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}

	fmt.Println("Successfully connected to UDP client!")
	UdpClient = conn
}
