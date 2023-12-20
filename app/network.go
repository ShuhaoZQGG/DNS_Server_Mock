package main

import (
	"net"
)

func readFromDnsServer(conn *net.UDPConn) ([]byte, *net.UDPAddr, error) {
	buf := make([]byte, 512)
	size, source, err := conn.ReadFromUDP(buf)
	if err != nil {
		return nil, source, err
	}

	receivedData := buf[:size]
	return receivedData, source, err
}

// processDnsRequest processes a DNS request and returns a DNS response packet.
func processDnsRequest(request *DNSPacket) *DNSPacket {
	// Process the request to create a response
	// This may involve looking up records, handling different types of queries, etc.
	// For simplicity, let's assume we're responding with a static record

	// Create the response components
	header := NewHeader(request.Header) // Modify as needed
	questions := request.Questions
	var answers []*Answer // Usually echoed back in the response
	for _, v := range questions {
		answer := NewAnswer(v.Name, 1, 1, 60, 4, []byte{8, 8, 8, 8})
		answers = append(answers, answer)
	}

	// Create the DNS response packet
	return NewDNSPacket(header, request.Questions, answers)
}
