package main

import (
	"encoding/binary"
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
