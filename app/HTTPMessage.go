package main

import (
	"bufio"
	"fmt"
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

func (httpMessage *HTTPMessage) writeHTTPMessage(writer *bufio.Writer) error {
	err := writeLine(writer, httpMessage.startLine.ToString())
	errorMessage := "failed to write http message"
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	for k, v := range httpMessage.headers {
		err = writeLine(writer, k+": "+v)
		if err != nil {
			return fmt.Errorf("%v: %v", errorMessage, err)
		}
	}

	err = writeLine(writer, "")
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	_, err = writer.WriteString(httpMessage.body)
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	return nil
}

func (httpMessage *HTTPMessage) printHTTPMessage() error {
	writer := bufio.NewWriter(os.Stdout)
	errorMessage := "failed to print http message"

	err := httpMessage.writeHTTPMessage(writer)
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	return nil
}

func printInputOutput(request *HTTPMessage, response *HTTPMessage) error {
	writer := bufio.NewWriter(os.Stdout)
	errorMessage := "failed to print http message input output"

	_, err := writer.WriteString("==============================\nRequest:\n")
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	err = request.writeHTTPMessage(writer)
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	_, err = writer.WriteString("==============================\nResponse:\n")
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	err = response.writeHTTPMessage(writer)
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	_, err = writer.WriteString("==============================\n")
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("%v: %v", errorMessage, err)
	}

	return nil
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

func parseBodyByLength(reader *bufio.Reader, length int) (string, error) {
	peek, err := reader.Peek(length)

	if err != nil {
		return "", fmt.Errorf("failed to parse body by length: %v", err)
	}

	return string(peek), nil
}

func parseBodyNoEOF(reader *bufio.Reader) string {
	return readLine(reader)
}

func (httpMessage *HTTPMessage) parseBody(reader *bufio.Reader) (string, error) {
	encodingType, hasEncoding := httpMessage.headers["transfer-encoding"]
	if hasEncoding && encodingType == "chunked" {
		return parseBodyByChunks(reader), nil
	}

	contentLength, hasLength := httpMessage.headers["content-length"]
	if hasLength {
		length, err := strconv.Atoi(contentLength)

		if err != nil {
			return "", fmt.Errorf("failed to parse body: %v", err)
		}

		return parseBodyByLength(reader, length)
	}

	return "", nil
}

func createEmptyHTTPMessage(startLine HTTPStartLine) *HTTPMessage {
	headers := map[string]string{}
	return &HTTPMessage{startLine: startLine, headers: headers, body: ""}
}

func (httpMessage *HTTPMessage) readHTTPMessage(reader *bufio.Reader) error {
	httpMessage.startLine = httpMessage.parseStartLine(reader)
	httpMessage.headers = parseHeaders(reader)
	body, err := httpMessage.parseBody(reader)

	if err != nil {
		return fmt.Errorf("failed to read http message: %v", err)
	}

	httpMessage.body = body

	return nil
}

func parseHTTPMessage(startLine HTTPStartLine, reader *bufio.Reader) (*HTTPMessage, error) {
	httpMessage := createEmptyHTTPMessage(startLine)
	err := httpMessage.readHTTPMessage(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http message: %v", err)
	}

	return httpMessage, nil
}
