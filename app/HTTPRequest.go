package main

import (
	"bufio"
	"strings"
)

type HTTPRequestLine struct {
	method  string
	path    string
	version string
}

func emptyHTTPRequestLine() HTTPRequestLine {
	return HTTPRequestLine{
		method:  "",
		path:    "",
		version: "",
	}
}

func (requestLine HTTPRequestLine) ToString() string {
	return strings.Join([]string{requestLine.method, requestLine.path, requestLine.version}, " ")
}

func (requestLine HTTPRequestLine) FromString(message string) HTTPStartLine {
	split := strings.Split(message, " ")

	method := ""
	path := ""
	version := ""

	if len(split) >= 3 {
		method = split[0]
		path = split[1]
		version = split[2]
	} else {
		warn("HTTP start line is empty")
	}

	return HTTPRequestLine{method: method, path: path, version: version}
}

func parseHTTPRequest(reader *bufio.Reader) (*HTTPMessage, error) {
	return parseHTTPMessage(emptyHTTPRequestLine(), reader)
}

func getHTTPRequestLine(httpMessage *HTTPMessage) (HTTPRequestLine, bool) {
	requestLine, ok := httpMessage.startLine.(HTTPRequestLine)
	return requestLine, ok
}
