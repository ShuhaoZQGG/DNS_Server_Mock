package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func createTestDNSPacket() []byte {
	var buf bytes.Buffer

	// Header
	buf.Write([]byte{0x00, 0x01}) // ID
	buf.Write([]byte{0x01, 0x00}) // Flags
	buf.Write([]byte{0x00, 0x02}) // QDCOUNT
	buf.Write([]byte{0x00, 0x00}) // ANCOUNT
	buf.Write([]byte{0x00, 0x00}) // NSCOUNT
	buf.Write([]byte{0x00, 0x00}) // ARCOUNT

	// Question 1: abc.longassdomainname.com
	buf.Write([]byte{0x03}) // Length of "abc"
	buf.WriteString("abc")
	buf.Write([]byte{0x10}) // Length of "longassdomainname"
	buf.WriteString("longassdomainname")
	buf.Write([]byte{0x03}) // Length of "com"
	buf.WriteString("com")
	buf.Write([]byte{0x00})       // Null terminator
	buf.Write([]byte{0x00, 0x01}) // QTYPE
	buf.Write([]byte{0x00, 0x01}) // QCLASS

	// Question 2: def.longassdomainname.com
	buf.Write([]byte{0x03}) // Length of "def"
	buf.WriteString("def")
	buf.Write([]byte{0x10}) // Length of "longassdomainname"
	buf.WriteString("longassdomainname")
	buf.Write([]byte{0x03}) // Length of "com"
	buf.WriteString("com")
	buf.Write([]byte{0x00})       // Null terminator
	buf.Write([]byte{0x00, 0x01}) // QTYPE
	buf.Write([]byte{0x00, 0x01}) // QCLASS

	return buf.Bytes()
}

func Test_parseDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "with google.com",
			args: args{
				domain: "google.com",
			},
			want: []byte("\x06google\x03com\x00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDomain() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Println("---Success: Test_parseDomain")
}

// TestParseDNSPacketWithOneQuestion tests the ParseDNSPacket function for accuracy.
func TestParseDNSPacketWithOneQuestion(t *testing.T) {
	// Example DNS packet data (you'll need to replace this with actual test data)
	data := []byte{162, 19, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 9, 115, 116, 97, 114, 45, 109, 105, 110, 105, 4, 99, 49, 48, 114, 8, 102, 97, 99, 101, 98, 111, 111, 107, 3, 99, 111, 109, 0, 0, 1, 0, 1}

	// Call ParseDNSPacket
	packet, err := ParseDNSPacket(data)
	if err != nil {
		t.Fatalf("Failed to parse DNS packet: %s", err)
	}

	// Validate the results (adjust these conditions based on your test data)
	const expectedID = 41491
	const expectedQuestionCount = 1
	const expectedQuestionName = "star-mini.c10r.facebook.com"
	if packet.Header.ID != expectedID {
		t.Errorf("Header ID mismatch: got %v, want %v", packet.Header.ID, expectedID)
	}
	if len(packet.Questions) != expectedQuestionCount {
		t.Errorf("Number of questions mismatch: got %d, want %d", len(packet.Questions), expectedQuestionCount)
	}

	if packet.Questions[0].Name != expectedQuestionName {
		t.Errorf("Question Name mismatch: got %s, want %s", packet.Questions[0].Name, expectedQuestionName)
	}
	fmt.Println("---Success: TestParseDNSPacketWithOneQuestion")
	// Add more checks as necessary for your test cases
}

// TestParseDNSPacketWithTwoQuestions tests the ParseDNSPacket function for accuracy.
func TestParseDNSPacketWithTwoQuestions(t *testing.T) {
	// Example DNS packet data (you'll need to replace this with actual test data)
	data := []byte{156, 123, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 3, 97, 98, 99, 17, 108, 111, 110, 103, 97, 115, 115, 100, 111, 109, 97, 105, 110, 110, 97, 109, 101, 3, 99, 111, 109, 0, 0, 1, 0, 1, 3, 100, 101, 102, 192, 16, 0, 1, 0, 1}

	// Call ParseDNSPacket
	packet, err := ParseDNSPacket(data)
	if err != nil {
		t.Fatalf("Failed to parse DNS packet: %s", err)
	}

	// Validate the results (adjust these conditions based on your test data)
	const expectedID = 40059
	const expectedQuestionCount = 2
	if packet.Header.ID != expectedID {
		t.Errorf("Header ID mismatch: got %v, want %v", packet.Header.ID, expectedID)
	}
	if len(packet.Questions) != expectedQuestionCount {
		t.Errorf("Number of questions mismatch: got %d, want %d", len(packet.Questions), expectedQuestionCount)
	}

	fmt.Println("---Success: TestParseDNSPacketWithTwoQuestions")
	// Add more checks as necessary for your test cases
}

// Additional test functions for ParseQuestion and ParseDomainName can be added here.
