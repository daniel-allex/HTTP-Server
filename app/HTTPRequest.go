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
		println("[WARNING] HTTP start line is empty")
	}

	return HTTPRequestLine{method: method, path: path, version: version}
}

func parseHTTPRequest(scanner *bufio.Scanner) *HTTPMessage {
	return parseHTTPMessage(emptyHTTPRequestLine(), scanner)
}

func getHTTPRequestLine(httpMessage *HTTPMessage) HTTPRequestLine {
	requestLine, ok := httpMessage.startLine.(HTTPRequestLine)
	exceptIfNotOk("Failed to cast HTTPStartLine to HTTPRequestLine", ok)

	return requestLine
}
