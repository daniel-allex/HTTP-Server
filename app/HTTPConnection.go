package main

import "net"

type HTTPConnection struct {
	conn *TCPConnection
}

func acceptHTTPConnection(listener *net.Listener) *HTTPConnection {
	return &HTTPConnection{acceptTCPConnection(listener)}
}

func (httpConn *HTTPConnection) nextRequest() *HTTPMessage {
	return parseHTTPRequest(httpConn.conn.Scanner())
}

func (httpConn *HTTPConnection) sendResponse(httpMessage *HTTPMessage) {
	writer := httpConn.conn.Writer()
	httpMessage.writeHTTPMessage(writer)
	err := writer.Flush()
	validateResult("Failed to flush response", err)
}

func (httpConn *HTTPConnection) Close() {
	httpConn.conn.Close()
}
