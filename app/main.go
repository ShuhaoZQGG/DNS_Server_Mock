package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		header := &Header{
			ID:      1234,
			QR:      true,
			QDCOUNT: 1,
		}

		question := &Question{
			Name:  "codecrafters.io",
			Type:  TypeNameToValue("A"),
			Class: TypeNameToValue("IN"),
		}

		dns := &DNS{
			header:   header,
			question: question,
		}

		dnsBytes := dns.ToBytes()

		_, err = udpConn.WriteToUDP(dnsBytes, source)

		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
