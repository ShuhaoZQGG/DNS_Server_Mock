package main

import (
	"net"
	"time"
)

func readFromDnsServer(conn *net.UDPConn) ([]byte, *net.UDPAddr, error) {
	buf := make([]byte, 512)
	size, source, err := conn.ReadFromUDP(buf)
	if err != nil {
		return nil, source, err
	}

	receivedData := buf[:size]
	return receivedData, source, err
}

func sendToResoler(packet *DNSPacket, connAddr string) (*DNSPacket, error) {
	conn, err := net.Dial("udp", connAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	timeoutDuration := 5 * time.Second
	conn.SetDeadline(time.Now().Add(timeoutDuration))

	bytes := packet.ToBytes()
	_, err = conn.Write(bytes)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	responsePacket, err := ParseDNSPacket(buffer[:n])
	if err != nil {
		return nil, err
	}

	return responsePacket, nil
}
