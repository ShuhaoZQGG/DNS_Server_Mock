package main

import "encoding/binary"

// DNS message header struct, for more information: https://www.oreilly.com/library/view/hands-on-network-programming/9781789349863/812dd5c5-0d22-4ccd-8faf-f339b416bb2e.xhtml

type Header struct {
	ID      uint16 // Packet Identifier 16 bits
	QR      bool   // Query/Response Indicator 1bit
	OPCODE  byte   // Operation Code 4 bits
	AA      bool   // Authoritative Answer 1 bit
	TC      bool   // Truncation 1 bit
	RD      bool   // Recusrion Desired 1 bit
	RA      bool   // Recursion Available 1 bit
	Z       byte   // Reserved 3 bits
	RCODE   byte   // Response Code 4 bits
	QDCOUNT uint16 // Question Count 16 bits
	ANCOUNT uint16 // Answer Record Count  16 bits
	NSCOUNT uint16 // Authority Record Count
	ARCOUNT uint16 // Additional Record Count
}

func (h *Header) ToBytes() []byte {
	// Create a byte slice of fixed size to hold the header data
	bytes := make([]byte, 12)
	// Convert the ID to bytes and store in the first two bytes
	binary.BigEndian.PutUint16(bytes[:2], h.ID)

	// Assemble the second and third bytes from various boolean flags and values.
	// Each part is shifted to its correct position in the byte.
	bytes[2] = BoolToByte(h.QR)<<7 | h.OPCODE<<3 | BoolToByte(h.AA)<<2 | BoolToByte(h.TC)<<1 | BoolToByte(h.RD)
	bytes[3] = BoolToByte(h.RA)<<7 | h.Z<<4 | h.RCODE
	// Convert the count fields to bytes and store them.
	binary.BigEndian.PutUint16(bytes[4:6], h.QDCOUNT)  // Query count
	binary.BigEndian.PutUint16(bytes[6:8], h.ANCOUNT)  // Answer count
	binary.BigEndian.PutUint16(bytes[8:10], h.NSCOUNT) // Authority records count
	binary.BigEndian.PutUint16(bytes[10:], h.ARCOUNT)  // Additional records count

	// Return the byte slice representation of the header
	return bytes
}

func BoolToByte(b bool) byte {
	if b == true {
		return 1
	}
	return 0
}
