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

func (responseLine HTTPResponseLine) setVersion(version string) HTTPResponseLine {
	responseLine.version = version

	return responseLine
}

func (responseLine HTTPResponseLine) setStatusCode(statusCode int) HTTPResponseLine {
	responseLine.statusCode = statusCode

	return responseLine
}

func (responseLine HTTPResponseLine) setStatusText(statusText string) HTTPResponseLine {
	responseLine.statusText = statusText

	return responseLine
}

func getHTTPResponseLine(httpMessage *HTTPMessage) HTTPResponseLine {
	responseLine := httpMessage.startLine.(HTTPResponseLine)
	return responseLine
}
