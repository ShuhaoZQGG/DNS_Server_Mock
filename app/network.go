package main

import (
	"net"
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
