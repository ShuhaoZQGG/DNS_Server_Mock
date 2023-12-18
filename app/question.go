package main

import (
	"encoding/binary"
)

type Question struct {
	Name  string
	Type  byte
	Class byte
}

func TypeNameToValue(s string) byte {
	switch s {
	case "A":
		return 1
	case "NS":
		return 2
	case "MD":
		return 3
	case "MF":
		return 4
	case "CNAME":
		return 5
	case "SOA":
		return 6
	case "MB":
		return 7
	case "MG":
		return 8
	case "MR":
		return 9
	case "NULL":
		return 10
	case "WKS":
		return 11
	case "PTR":
		return 12
	case "HINFO":
		return 13
	case "MINFO":
		return 14
	case "MX":
		return 15
	case "TXT":
		return 16
	default:
		return 0
	}
}

func ClassNameToValue(s string) byte {
	switch s {
	case "IN":
		return 1
	case "CS":
		return 2
	case "CH":
		return 3
	case "HS":
		return 4
	default:
		return 0
	}
}

func ByteToBigEndianInt(value byte) []byte {
	var valueUInt16 uint16 = uint16(value)
	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, valueUInt16)
	return result
}

func (q *Question) ToBytes() []byte {
	var question []byte
	var domain []byte = ParseDomain(q.Name)
	var types []byte = ByteToBigEndianInt(q.Type)
	var class []byte = ByteToBigEndianInt(q.Class)
	question = append(domain, types...)
	question = append(question, class...)
	return question
}
