package main

import (
	"bufio"
	"net"
)

type TCPConnection struct {
	reader *bufio.Reader
	writer *bufio.Writer
	conn   *net.Conn
}

func createTCPConnection(conn *net.Conn) *TCPConnection {
	reader := bufio.NewReader(*conn)
	writer := bufio.NewWriter(*conn)

	return &TCPConnection{reader: reader, writer: writer, conn: conn}
}

func listenTCPConnection(address string) *net.Listener {
	listener, err := net.Listen("tcp", address)
	validateResult("Failed to bind to port "+address, err)
	return &listener
}

func acceptTCPConnection(listener *net.Listener) *TCPConnection {
	conn, err := (*listener).Accept()
	validateResult("Error accepting connection", err)
	return createTCPConnection(&conn)
}

func (conn *TCPConnection) Close() {
	err := (*conn.conn).Close()
	validateResult("Failed to close connection", err)
}

func (conn *TCPConnection) Reader() *bufio.Reader {
	return conn.reader
}

func (conn *TCPConnection) Writer() *bufio.Writer {
	return conn.writer
}
