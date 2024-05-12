package main

import (
	"bufio"
	"os"
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

func (httpMessage *HTTPMessage) printHTTPMessage() {
	writer := bufio.NewWriter(os.Stdout)
	httpMessage.writeHTTPMessage(writer)
	err := writer.Flush()
	validateResult("Failed to flush writer to STDOUT", err)
}

func printInputOutput(request *HTTPMessage, response *HTTPMessage) {
	writer := bufio.NewWriter(os.Stdout)
	write(writer, "==============================\n")
	write(writer, "Request:\n")
	request.writeHTTPMessage(writer)
	write(writer, "------------------------------\n")
	write(writer, "Response:\n")
	response.writeHTTPMessage(writer)
	write(writer, "==============================\n")

	err := writer.Flush()
	validateResult("Failed to flush writer to STDOUT", err)

}

func (httpMessage *HTTPMessage) parseStartLine(reader *bufio.Reader) HTTPStartLine {
	firstLine := readLine(reader)
	return httpMessage.startLine.FromString(firstLine)
}

func parseHeaders(reader *bufio.Reader) map[string]string {
	headers := map[string]string{}

	nextLine := readLine(reader)
	for nextLine != "" {
		key, val, _ := strings.Cut(nextLine, ": ")
		headers[strings.ToLower(key)] = val

		nextLine = readLine(reader)
	}

	return headers
}

func parseBodyByChunks(reader *bufio.Reader) string {
	var sb strings.Builder
	for nextBytes := readLine(reader); nextBytes != "0"; {
		sb.WriteString(readLine(reader))
	}

	readLine(reader)

	return sb.String()
}

func parseBodyByLength(reader *bufio.Reader, length int) string {
	peek, err := reader.Peek(length)
	validateResult("Failed to peek reader when parsing body by length", err)

	return string(peek)
}

func parseBodyNoEOF(reader *bufio.Reader) string {
	return readLine(reader)
}

func (httpMessage *HTTPMessage) parseBody(reader *bufio.Reader) string {
	encodingType, hasEncoding := httpMessage.headers["transfer-encoding"]
	if hasEncoding && encodingType == "chunked" {
		return parseBodyByChunks(reader)
	}

	contentLength, hasLength := httpMessage.headers["content-length"]
	if hasLength {
		length, err := strconv.Atoi(contentLength)
		validateResult("Failed to convert content length to int", err)
		return parseBodyByLength(reader, length)
	}

	return ""
}

func createEmptyHTTPMessage(startLine HTTPStartLine) *HTTPMessage {
	headers := map[string]string{}
	return &HTTPMessage{startLine: startLine, headers: headers, body: ""}
}

func (httpMessage *HTTPMessage) readHTTPMessage(reader *bufio.Reader) {
	httpMessage.startLine = httpMessage.parseStartLine(reader)
	httpMessage.headers = parseHeaders(reader)
	httpMessage.body = httpMessage.parseBody(reader)
}

func parseHTTPMessage(startLine HTTPStartLine, reader *bufio.Reader) *HTTPMessage {
	httpMessage := createEmptyHTTPMessage(startLine)
	httpMessage.readHTTPMessage(reader)

	return httpMessage
}
