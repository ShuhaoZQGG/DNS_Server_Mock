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

	for {
		packet, source, err := readFromDnsServer(udpConn)
		if err != nil {
			fmt.Printf("Error reading DNS packet: %s\n", err)
			break
		}

		receivedHeader, err := ParseHeader(packet)

		if err != nil {
			fmt.Printf("Error parsing header: %s\n", err)
			break
		}

		header := NewHeader(receivedHeader)

		question := &Question{
			Name:  "codecrafters.io",
			Type:  TypeNameToValue("A"),
			Class: ClassNameToValue("IN"),
		}

		answer := &Answer{
			Name:   "codecrafters.io",
			Type:   TypeNameToValue("A"),
			Class:  ClassNameToValue("IN"),
			TTL:    60,
			Length: 4,
			Data:   "8.8.8.8",
		}

		dns := &DNS{
			Header:   header,
			Question: question,
			Answer:   answer,
		}

		dnsBytes := dns.ToBytes()

		_, err = udpConn.WriteToUDP(dnsBytes, source)

		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
