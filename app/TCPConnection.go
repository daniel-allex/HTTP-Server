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

func listenTCPConnection(address string) (*net.Listener, error) {
	listener, err := net.Listen("tcp", address)
	return &listener, err
}

func acceptTCPConnection(listener *net.Listener) (*TCPConnection, error) {
	conn, err := (*listener).Accept()

	return createTCPConnection(&conn), err
}

func (conn *TCPConnection) Close() error {
	err := (*conn.conn).Close()

	return err
}

func (conn *TCPConnection) Reader() *bufio.Reader {
	return conn.reader
}

func (conn *TCPConnection) Writer() *bufio.Writer {
	return conn.writer
}
