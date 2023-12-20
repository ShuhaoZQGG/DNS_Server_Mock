package main

type DNSPacket struct {
	Header    *Header
	Questions []*Question
	Answers   []*Answer
}

func NewDNSPacket(header *Header, questions []*Question, answers []*Answer) *DNSPacket {
	return &DNSPacket{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}
}

func (packet *DNSPacket) ToBytes() []byte {
	var response []byte
	response = append(response, packet.Header.ToBytes()...)
	response = append(response, flattenQuestions(packet.Questions)...)
	response = append(response, flattenAnswers(packet.Answers)...)
	return response
}

type DNSPackets []*DNSPacket

func (packets DNSPackets) ToBytes() []byte {
	var response []byte

	for _, packet := range packets {
		response = append(response, packet.Header.ToBytes()...)
		response = append(response, flattenQuestions(packet.Questions)...)
		response = append(response, flattenAnswers(packet.Answers)...)
	}

	return response
}

// ParseDNSPacket parses the DNS packet and extracts the header, question, and answer.
func ParseDNSPacket(data []byte) (*DNSPacket, error) {
	// fmt.Println("data in bytes", data)
	// Assuming the header is already parsed and part of the data
	header, length, err := ParseHeader(data)
	if err != nil {
		return nil, err
	}

	var questions []*Question

	offset := length // Adjusting the offset to start parsing questions

	for i := 0; i < int(header.QDCOUNT); i++ {
		question, n, err := ParseQuestion(data, offset)
		if err != nil {
			return nil, err
		}
		offset = n
		questions = append(questions, question)
	}

	var answers []*Answer
	for i := 0; i < int(header.ANCOUNT); i++ {
		answer, n, err := ParseAnswer(data, offset)
		if err != nil {
			return nil, err
		}
		offset = n
		answers = append(answers, answer)
	}

	return &DNSPacket{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}, nil
}

func SplitDNSPacket(dnsPacket *DNSPacket) (dnsPackets []*DNSPacket) {
	if dnsPacket.Header.QDCOUNT <= 1 {
		return append(dnsPackets, dnsPacket)
	}

	for _, v := range dnsPacket.Questions {
		var questions []*Question
		header := dnsPacket.Header
		header.QDCOUNT = 1
		newDnsPacket := &DNSPacket{
			Header:    header,
			Questions: append(questions, v),
		}
		dnsPackets = append(dnsPackets, newDnsPacket)
	}

	return
}

func CombineDNSPackets(dnsPackets []*DNSPacket) *DNSPacket {
	header := dnsPackets[0].Header
	header.QDCOUNT = uint16(len(dnsPackets))
	header.ANCOUNT = uint16(len(dnsPackets))
	var questions []*Question
	var answers []*Answer

	// Assuming the Header, Questions, and Answers are properly defined in DNSPacket
	for _, response := range dnsPackets {
		questions = append(questions, response.Questions...)
		answers = append(answers, response.Answers...)
	}

	// Create a new DNS packet with combined questions and answers
	return &DNSPacket{
		Header:    header, // You'll need to set the appropriate header fields
		Questions: questions,
		Answers:   answers,
	}
}
