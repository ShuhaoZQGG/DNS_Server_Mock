package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseDomain(domain string) []byte {
	var response []byte
	subDomains := strings.Split(domain, ".")
	for _, subDomain := range subDomains {
		response = append(response, byte(len(subDomain)))
		response = append(response, subDomain...)
	}
	response = append(response, '\x00')
	return response
}

func ParseDomainName(data []byte) (string, int) {
	var nameParts []string
	offset := 0

	for {
		if offset >= len(data) {
			return "", offset
		}

		length := int(data[offset])
		if length == 0 {
			break
		}
		offset++

		if offset+length > len(data) {
			return "", offset
		}

		// Append the current part of the domain name
		nameParts = append(nameParts, string(data[offset:offset+length]))
		offset += length
	}

	// Join the parts of the domain name with "." to form the full domain name
	return strings.Join(nameParts, "."), offset + 1
}

func ParseIP(ip string) []byte {
	var response []byte
	elements := strings.Split(ip, ".")
	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
		}
		response = append(response, byte(num))
	}
	return response
}

func TypeNameToValue(s string) uint16 {
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

func ClassNameToValue(s string) uint16 {
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
