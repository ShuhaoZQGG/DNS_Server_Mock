package main

import (
	"fmt"
)

type DNSPacket struct {
	Header   *Header
	Question *Question
	Answer   *Answer
}

func NewDNSPacket(header *Header, question *Question, answer *Answer) *DNSPacket {
	return &DNSPacket{
		Header:   header,
		Question: question,
		Answer:   answer,
	}
}

func (packet *DNSPacket) ToBytes() []byte {
	var response []byte
	response = append(response, packet.Header.ToBytes()...)
	response = append(response, packet.Question.ToBytes()...)
	response = append(response, packet.Answer.ToBytes()...)
	return response
}

// ParseDNSPacket parses the DNS packet and extracts the header, question, and answer.
func ParseDNSPacket(data []byte) (*DNSPacket, error) {
	header, off, err := ParseHeader(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing header: %w", err)
	}

	question, off, err := ParseQuestion(data[off:])
	if err != nil {
		return nil, fmt.Errorf("error parsing question: %w", err)
	}

	return &DNSPacket{
		Header:   header,
		Question: question,
	}, nil
}
