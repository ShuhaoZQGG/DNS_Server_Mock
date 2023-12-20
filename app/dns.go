package main

import "fmt"

type DNSPacket struct {
	Header    *Header
	Questions []*Question
	Answers   []*Answer
}

func NewDNSPacket(header *Header, questions []*Question, answers []*Answer) *DNSPacket {
	return &DNSPacket{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}
}

func (packet *DNSPacket) ToBytes() []byte {
	var response []byte
	response = append(response, packet.Header.ToBytes()...)
	response = append(response, flattenQuestions(packet.Questions)...)
	response = append(response, flattenAnswers(packet.Answers)...)
	return response
}

// ParseDNSPacket parses the DNS packet and extracts the header, question, and answer.
func ParseDNSPacket(data []byte) (*DNSPacket, error) {
	fmt.Println("data in bytes", data)
	// Assuming the header is already parsed and part of the data
	header, length, err := ParseHeader(data)
	if err != nil {
		return nil, err
	}

	var questions []*Question

	offset := length // Adjusting the offset to start parsing questions

	for i := 0; i < int(header.QDCOUNT); i++ {
		question, n, err := ParseQuestion(data, offset)
		if err != nil {
			return nil, err
		}
		offset = n
		questions = append(questions, question)
	}

	return &DNSPacket{
		Header:    header,
		Questions: questions,
	}, nil
}

func flattenQuestions(questions []*Question) []byte {
	var bytes []byte
	for _, q := range questions {
		bytes = append(bytes, q.ToBytes()...)
	}
	return bytes
}

func flattenAnswers(answers []*Answer) []byte {
	var bytes []byte
	for _, a := range answers {
		bytes = append(bytes, a.ToBytes()...)
	}
	return bytes
}
