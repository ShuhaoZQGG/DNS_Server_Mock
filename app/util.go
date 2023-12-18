package main

import (
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
