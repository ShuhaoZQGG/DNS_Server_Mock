package main

import (
	"encoding/binary"
	"fmt"
)

type Answer struct {
	Name   string
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   []byte
}

func NewAnswer(name string, rtype uint16, class uint16, ttl uint32, length uint16, data []byte) *Answer {
	return &Answer{
		Name:   name,
		Type:   rtype,
		Class:  class,
		TTL:    ttl,
		Length: length,
		Data:   data,
	}
}

func (a *Answer) ToBytes() []byte {
	var answer []byte
	nameBytes := ParseDomain(a.Name)
	answer = append(answer, nameBytes...)
	answer = append(answer, uint16ToBytes(a.Type)...)
	answer = append(answer, uint16ToBytes(a.Class)...)
	answer = append(answer, uint32ToBytes(a.TTL)...)
	answer = append(answer, uint16ToBytes(a.Length)...)
	answer = append(answer, a.Data...)
	return answer
}

func uint16ToBytes(value uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, value)
	return bytes
}

func uint32ToBytes(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, value)
	return bytes
}

func ParseAnswer(data []byte, initialOffset int) (*Answer, int, error) {
	name, offset := ParseDomainName(data, initialOffset)

	// Make sure there's enough data for the fixed-length fields of the answer (Type, Class, TTL, Length)
	if offset+10 > len(data) {
		return nil, 0, fmt.Errorf("answer section is too short, data length: %d, needed: %d", len(data), offset+10)
	}

	atype := binary.BigEndian.Uint16(data[offset : offset+2])
	aclass := binary.BigEndian.Uint16(data[offset+2 : offset+4])
	ttl := binary.BigEndian.Uint32(data[offset+4 : offset+8])
	dataLength := binary.BigEndian.Uint16(data[offset+8 : offset+10])

	// Check if there's enough data for the Data field
	if offset+10+int(dataLength) > len(data) {
		return nil, 0, fmt.Errorf("answer data is too short, data length: %d, needed: %d", len(data), offset+10+int(dataLength))
	}

	answerData := make([]byte, dataLength)
	copy(answerData, data[offset+10:offset+10+int(dataLength)])

	answer := &Answer{
		Name:   name,
		Type:   atype,
		Class:  aclass,
		TTL:    ttl,
		Length: dataLength,
		Data:   answerData,
	}

	return answer, offset + 10 + int(dataLength), nil
}

func flattenAnswers(answers []*Answer) []byte {
	var bytes []byte
	for _, a := range answers {
		bytes = append(bytes, a.ToBytes()...)
	}
	return bytes
}
