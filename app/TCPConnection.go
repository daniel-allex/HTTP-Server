package main

import (
	"bufio"
	"net"
)

type TCPConnection struct {
	reader bufio.Reader
	writer bufio.Writer
	conn   net.Conn
}

func createTCPConnection(conn net.Conn) TCPConnection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	return TCPConnection{reader: *reader, writer: *writer, conn: conn}
}

func connectTCP(protocol string, address string) TCPConnection {
	l, err := net.Listen(protocol, address)
	validateResult("Failed to bind to port "+address, err)

	conn, err := l.Accept()
	validateResult("Error accepting connection", err)

	return createTCPConnection(conn)
}

func (conn *TCPConnection) Close() {
	err := conn.conn.Close()
	validateResult("Failed to close connection", err)
}
