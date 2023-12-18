package main

type DNS struct {
	header   *Header
	question *Question
}

func (dns *DNS) ToBytes() []byte {
	var response []byte
	header := dns.header.ToBytes()
	question := dns.question.ToBytes()
	response = append(header, question...)
	return response
}
