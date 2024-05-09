package main

import (
	"strconv"
	"strings"
)

type HTTPResponseLine struct {
	version    string
	statusCode int
	statusText string
}

func emptyHTTPResponseLine() HTTPResponseLine {
	return HTTPResponseLine{
		version:    "",
		statusCode: 0,
		statusText: "",
	}
}

func (responseLine HTTPResponseLine) ToString() string {
	statusCode := strconv.Itoa(responseLine.statusCode)
	return strings.Join([]string{responseLine.version, statusCode, responseLine.statusText}, " ")
}

func (responseLine HTTPResponseLine) FromString(message string) HTTPStartLine {
	split := strings.Split(message, " ")
	version, statusCode, statusText := split[0], split[1], split[2]
	code, err := strconv.Atoi(statusCode)

	validateResult("Failed to convert status code from string to int", err)

	return HTTPResponseLine{version: version, statusCode: code, statusText: statusText}
}

func getHTTPResponseLine(httpMessage HTTPMessage) *HTTPResponseLine {
	responseLine, ok := httpMessage.startLine.(HTTPResponseLine)
	exceptIfNotOk("Failed to cast HTTPStartLine to HTTPResponseLine", ok)

	return &responseLine
}
