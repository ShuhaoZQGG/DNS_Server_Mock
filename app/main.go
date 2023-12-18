package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting DNS server...")

	// Resolve the address to listen on
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	// Listen on the resolved UDP address
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("DNS server is listening...")

	for {
		// Read a packet from the DNS server
		packet, source, err := readFromDnsServer(udpConn)
		if err != nil {
			fmt.Printf("Error reading DNS packet: %s\n", err)
			continue
		}

		// Parse the received DNS packet
		dnsPacket, err := ParseDNSPacket(packet)
		if err != nil {
			fmt.Printf("Error parsing DNS packet: %s\n", err)
			continue
		}

		// Process the DNS request and prepare the response
		dnsReply := processDnsRequest(dnsPacket)
		if dnsReply == nil {
			fmt.Println("Failed to process DNS request")
			continue
		}

		// Convert the DNS response to bytes
		dnsBytes := dnsReply.ToBytes()

		// Send the response back to the source
		_, err = udpConn.WriteToUDP(dnsBytes, source)
		if err != nil {
			fmt.Println("Failed to send DNS response:", err)
		}
	}
}
