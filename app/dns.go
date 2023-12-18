package main

type DNS struct {
	Header   *Header
	Question *Question
	Answer   *Answer
}

func (dns *DNS) ToBytes() []byte {
	var response []byte
	header := dns.Header.ToBytes()
	question := dns.Question.ToBytes()
	answer := dns.Answer.ToBytes()
	response = append(header, question...)
	response = append(response, answer...)
	return response
}
