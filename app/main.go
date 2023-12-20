package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting DNS server...")

	address := flag.String("resolver", "localhost:3000", "ip:port")
	flag.Parse()

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

		timeout := 5 * time.Second

		// 1. parse packet into dnsPacket
		// Parse the received DNS packet
		dnsPacket, err := ParseDNSPacket(packet)
		if err != nil {
			fmt.Printf("Error parsing DNS packet: %s\n", err)
			continue
		}
		// 2. if dnsPacket.QDCount > 1, transfer dnsPacket { Header, []Questions} into []dnsPacket { Header, Question}
		dnsPackets := DNSPackets(SplitDNSPacket(dnsPacket))

		// 3. parse dnsPacket back into []bytes
		responses := make(chan *DNSPacket, len(dnsPackets))

		// Send packets to resolver in parallel
		for _, pkt := range dnsPackets {
			go func(packet *DNSPacket) {
				response, err := sendToResoler(packet, *address)
				if err != nil {
					fmt.Println("Error sending to resolver", err)
					return
				}
				responses <- response
			}(pkt)
		}

		// get dns packet responses back
		var collectedResponses []*DNSPacket
		for i := 0; i < len(dnsPackets); i++ {
			select {
			case response := <-responses:
				collectedResponses = append(collectedResponses, response)
			case <-time.After(timeout):
				log.Println("Timeout waiting for responses")
				break
			}
		}

		combinedPacket := CombineDNSPackets(collectedResponses)

		// Convert the DNS response to bytes
		dnsBytes := combinedPacket.ToBytes()
		// Send the response back to the source
		_, err = udpConn.WriteToUDP(dnsBytes, source)
		if err != nil {
			fmt.Println("Failed to send DNS response:", err)
		}

		if err != nil {
			fmt.Printf("Error reading DNS packet: %s\n", err)
			continue
		}
	}
}
