package main

import (
	"fmt"
	"net"
)

type HTTPConnection struct {
	conn *TCPConnection
}

func acceptHTTPConnection(listener *net.Listener) (*HTTPConnection, error) {
	connection, err := acceptTCPConnection(listener)
	if err != nil {
		return nil, err
	}
	return &HTTPConnection{connection}, nil
}

func (httpConn *HTTPConnection) nextRequest() (*HTTPMessage, error) {
	return parseHTTPRequest(httpConn.conn.Reader())
}

func (httpConn *HTTPConnection) sendResponse(httpMessage *HTTPMessage) error {
	writer := httpConn.conn.Writer()
	err := httpMessage.writeHTTPMessage(writer)
	if err != nil {
		return fmt.Errorf("failed to send http response: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to send http response: %v", err)
	}

	return nil
}

func (httpConn *HTTPConnection) Close() error {
	return httpConn.conn.Close()
}
