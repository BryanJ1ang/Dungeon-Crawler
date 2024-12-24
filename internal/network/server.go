package network

import (
	"fmt"
	"net"
)

func StartGameServer(port string) {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:" + port)

	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		fmt.Println("Error listening on UDP address:", err)
		return
	}

	defer conn.Close()

	fmt.Printf("UDP server is running on %s\n", "127.0.0.1")

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Handle the incoming client data
		HandleClientData(buffer[:n], addr)
	} 
}

func HandleClientData(data []byte, addr *net.UDPAddr) {
	MovementPacket := &MovementPacket{}
	MovementPacket.Decode(data)	
	fmt.Println("Character ID:", MovementPacket.CharacterID)
	fmt.Println("Direction:", MovementPacket.Direction)
}
