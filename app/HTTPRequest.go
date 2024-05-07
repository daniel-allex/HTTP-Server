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
	method, path, version := split[0], split[1], split[2]

	return HTTPRequestLine{method: method, path: path, version: version}
}

func parseHTTPRequest(reader bufio.Reader) HTTPMessage {
	return parseHTTPMessage(emptyHTTPRequestLine(), reader)
}
