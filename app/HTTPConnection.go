package main

import "strings"

type HTTPMessageContent struct {
	method          string
	path            string
	version         string
	headerHost      string
	headerUserAgent string
}

type HTTPConnection struct {
	conn TCPConnection
}

func connectHTTP(protocol string, address string) HTTPConnection {
	return HTTPConnection{connectTCP(protocol, address)}
}

func (HTTPConn *HTTPConnection) write(message string) {
	HTTPConn.conn.write(message)
}

func (HTTPConn *HTTPConnection) readLine() string {
	return strings.Trim(HTTPConn.conn.readLine(), "\r")
}

func (HTTPConn *HTTPConnection) WriteSuccessResponse(message string) {
	HTTPConn.write("HTTP/1.1 200 " + message + "\r\r\n\n")
}

func (HTTPConn *HTTPConnection) WriteFailedResponse(message string) {
	HTTPConn.write("HTTP/1.1 404 " + message + "\r\r\n\n")
}

func (HTTPConn *HTTPConnection) Close() {
	HTTPConn.conn.Close()
}

/*
Extracts the first line

Returns:
method, path, version
*/
func extractStartLine(firstLine string) []string {
	return strings.Split(firstLine, " ")
}

func (HTTPConn *HTTPConnection) nextRequest() HTTPMessageContent {
	firstLineRes := extractStartLine(HTTPConn.readLine())

	method, path, version := firstLineRes[0], firstLineRes[1], firstLineRes[2]

	return HTTPMessageContent{
		method:  method,
		path:    path,
		version: version,
	}
}
