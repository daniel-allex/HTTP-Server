package main

import (
	"bufio"
	"strconv"
	"strings"
)

type HTTPStartLine interface {
	ToString() string
	FromString(message string) HTTPStartLine
}

type HTTPMessage struct {
	startLine HTTPStartLine
	headers   map[string]string
	body      string
}

func (httpMessage *HTTPMessage) writeHTTPMessage(writer *bufio.Writer) {
	writeLine(writer, httpMessage.startLine.ToString())

	for k, v := range httpMessage.headers {
		writeLine(writer, k+": "+v)
	}

	writeLine(writer, "")
	write(writer, httpMessage.body)
}

func (httpMessage *HTTPMessage) parseStartLine(scanner *bufio.Scanner) HTTPStartLine {
	firstLine := readLine(scanner)
	return httpMessage.startLine.FromString(firstLine)
}

func parseHeaders(scanner *bufio.Scanner) map[string]string {
	headers := map[string]string{}

	nextLine := readLine(scanner)
	for nextLine != "" {
		key, val, _ := strings.Cut(nextLine, ": ")
		headers[strings.ToLower(key)] = val

		nextLine = readLine(scanner)
	}

	return headers
}

func parseBodyByChunks(scanner *bufio.Scanner) string {
	var sb strings.Builder
	for nextBytes := readLine(scanner); nextBytes != "0"; {
		sb.WriteString(readLine(scanner))
	}

	readLine(scanner)

	return sb.String()
}

func parseBodyByLength(scanner *bufio.Scanner, length int) string {
	return (readLine(scanner))[:length]
}

func parseBodyNoEOF(scanner *bufio.Scanner) string {
	return readLine(scanner)
}

func (httpMessage *HTTPMessage) parseBody(scanner *bufio.Scanner) string {
	encodingType, hasEncoding := httpMessage.headers["Transfer-Encoding"]
	if hasEncoding && encodingType == "chunked" {
		return parseBodyByChunks(scanner)
	}

	contentLength, hasLength := httpMessage.headers["ContentLength"]
	if hasLength {
		length, err := strconv.Atoi(contentLength)
		validateResult("Failed to convert content length to int", err)
		return parseBodyByLength(scanner, length)
	}

	return parseBodyNoEOF(scanner)
}

func createEmptyHTTPMessage(startLine HTTPStartLine) *HTTPMessage {
	headers := map[string]string{}
	return &HTTPMessage{startLine: startLine, headers: headers, body: ""}
}

func (httpMessage *HTTPMessage) scanHTTPMessage(scanner *bufio.Scanner) {
	httpMessage.startLine = httpMessage.parseStartLine(scanner)
	httpMessage.headers = parseHeaders(scanner)
	// httpMessage.body = httpMessage.parseBody(scanner)
}

func parseHTTPMessage(startLine HTTPStartLine, scanner *bufio.Scanner) *HTTPMessage {
	httpMessage := createEmptyHTTPMessage(startLine)
	httpMessage.scanHTTPMessage(scanner)

	return httpMessage
}
