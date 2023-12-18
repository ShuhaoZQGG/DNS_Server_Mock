package main

import (
	"encoding/binary"
	"fmt"
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func NewQuestion(name string, qtype uint16, class uint16) *Question {
	return &Question{
		Name:  name,
		Type:  qtype,
		Class: class,
	}
}

func (q *Question) ToBytes() []byte {
	var question []byte
	question = append(question, ParseDomain(q.Name)...)
	question = append(question, uint16ToBytes(q.Type)...)
	question = append(question, uint16ToBytes(q.Class)...)
	return question
}

// ParseQuestion parses the question part of a DNS packet.
func ParseQuestion(data []byte) (*Question, int, error) {
	name, offset := ParseDomainName(data)
	if offset+4 > len(data) {
		return nil, 0, fmt.Errorf("question section is too short")
	}

	qtype := binary.BigEndian.Uint16(data[offset : offset+2])
	qclass := binary.BigEndian.Uint16(data[offset+2 : offset+4])

	question := &Question{
		Name:  name,
		Type:  qtype,
		Class: qclass,
	}

	return question, offset + 4, nil
}
