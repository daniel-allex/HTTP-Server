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

func (httpMessage *HTTPMessage) writeHTTPMessage(writer bufio.Writer) {
	writeLine(writer, httpMessage.startLine.ToString())

	for k, v := range httpMessage.headers {
		writeLine(writer, k+": "+v)
	}

	writeLine(writer, "")
	write(writer, httpMessage.body)
}

func (httpMessage *HTTPMessage) parseStartLine(reader bufio.Reader) HTTPStartLine {
	return httpMessage.startLine.FromString(readLine(reader))
}

func parseHeaders(reader bufio.Reader) map[string]string {
	var headers map[string]string

	nextLine := readLine(reader)
	for nextLine != "" {
		key, val, _ := strings.Cut(nextLine, ": ")
		headers[key] = val

		nextLine = readLine(reader)
	}

	return headers
}

func parseBodyByChunks(reader bufio.Reader) string {
	nextBytes := readLine(reader)

	var sb strings.Builder
	for nextBytes != "0" {
		sb.WriteString(readLine(reader))
		nextBytes = readLine(reader)
	}

	// Discard EOF
	readLine(reader)

	return sb.String()
}

func parseBodyByLength(reader bufio.Reader, length int) string {
	bytes, err := reader.Peek(length)
	validateResult("Failed to read body by content length", err)

	return string(bytes)
}

func parseBodyNoEOF(reader bufio.Reader) string {
	var buffer []byte
	_, err := reader.Read(buffer)
	validateResult("Failed to read body without EOF", err)

	return string(buffer)
}

func (httpMessage *HTTPMessage) parseBody(reader bufio.Reader) string {
	encodingType, hasEncoding := httpMessage.headers["Transfer-Encoding"]
	if hasEncoding && encodingType == "chunked" {
		return parseBodyByChunks(reader)
	}

	contentLength, hasLength := httpMessage.headers["ContentLength"]
	if hasLength {
		length, err := strconv.Atoi(contentLength)
		validateResult("Failed to convert content length to int", err)
		return parseBodyByLength(reader, length)
	}

	return parseBodyNoEOF(reader)
}

func createEmptyHTTPMessage(startLine HTTPStartLine) HTTPMessage {
	var headers map[string]string
	return HTTPMessage{startLine: startLine, headers: headers, body: ""}
}

func (httpMessage *HTTPMessage) readHTTPMessage(reader bufio.Reader) {
	httpMessage.startLine = httpMessage.parseStartLine(reader)
	httpMessage.headers = parseHeaders(reader)
	httpMessage.body = httpMessage.parseBody(reader)
}

func parseHTTPMessage(startLine HTTPStartLine, reader bufio.Reader) HTTPMessage {
	httpMessage := createEmptyHTTPMessage(startLine)
	httpMessage.readHTTPMessage(reader)

	return httpMessage
}
