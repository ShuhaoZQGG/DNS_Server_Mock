package main

import "encoding/binary"

type Answer struct {
	Name   string
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   string
}

func (a *Answer) ToBytes() []byte {
	var answer []byte
	var nameBytes []byte = ParseDomain(a.Name)
	answer = append(answer, nameBytes...)

	typeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(typeBytes, a.Type)
	answer = append(answer, typeBytes...)

	classBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(classBytes, a.Class)
	answer = append(answer, classBytes...)

	ttlBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(ttlBytes, a.TTL)
	answer = append(answer, ttlBytes...)

	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, a.Length)
	answer = append(answer, lengthBytes...)

	var dataBytes []byte = ParseIP(a.Data)
	answer = append(answer, dataBytes...)

	return answer
}
