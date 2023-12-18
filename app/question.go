package main

import "encoding/binary"

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (q *Question) ToBytes() []byte {
	var question []byte
	var domain []byte = ParseDomain(q.Name)
	typeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(typeBytes, q.Type)
	classBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(classBytes, q.Class)
	question = append(question, domain...)
	question = append(domain, typeBytes...)
	question = append(question, classBytes...)
	return question
}
