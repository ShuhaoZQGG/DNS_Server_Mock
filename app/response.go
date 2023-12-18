package main

type DNS struct {
	header   *Header
	question *Question
}

func (dns *DNS) ToBytes() []byte {
	var response []byte
	header := dns.header.ToBytes()
	question := dns.header.ToBytes()
	response = append(header, question...)
	return response
}
