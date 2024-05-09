package main

import (
	"bufio"
	"net"
)

type TCPConnection struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
	conn    *net.Conn
}

func createTCPConnection(conn *net.Conn) *TCPConnection {
	connection := *conn
	scanner := bufio.NewScanner(connection)
	writer := bufio.NewWriter(connection)

	return &TCPConnection{scanner: scanner, writer: writer, conn: conn}
}

func connectTCP(protocol string, address string) *TCPConnection {
	l, err := net.Listen(protocol, address)
	validateResult("Failed to bind to port "+address, err)

	conn, err := l.Accept()
	validateResult("Error accepting connection", err)

	return createTCPConnection(&conn)
}

func (conn *TCPConnection) Close() {
	connection := *conn.conn
	err := connection.Close()
	validateResult("Failed to close connection", err)
}

func (conn *TCPConnection) Scanner() *bufio.Scanner {
	return conn.scanner
}

func (conn *TCPConnection) Writer() *bufio.Writer {
	return conn.writer
}
