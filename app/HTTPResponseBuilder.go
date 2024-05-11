package main

import (
	"strconv"
	"strings"
)

type HTTPResponseBuilder struct {
	httpMessage *HTTPMessage
}

func createHTTPResponseBuilder() *HTTPResponseBuilder {
	return &HTTPResponseBuilder{createEmptyHTTPMessage(emptyHTTPResponseLine())}
}

func (builder *HTTPResponseBuilder) setVersion(version string) *HTTPResponseBuilder {
	builder.httpMessage.startLine = getHTTPResponseLine(builder.httpMessage).setVersion(version)

	return builder
}

func (builder *HTTPResponseBuilder) setStatusCode(statusCode int) *HTTPResponseBuilder {
	builder.httpMessage.startLine = getHTTPResponseLine(builder.httpMessage).setStatusCode(statusCode)

	return builder
}

func (builder *HTTPResponseBuilder) setStatusText(statusText string) *HTTPResponseBuilder {
	builder.httpMessage.startLine = getHTTPResponseLine(builder.httpMessage).setStatusText(statusText)

	return builder
}

func (builder *HTTPResponseBuilder) setHeader(key string, val string) *HTTPResponseBuilder {
	builder.httpMessage.headers[strings.ToLower(key)] = val

	return builder
}

func (builder *HTTPResponseBuilder) setBody(message string) *HTTPResponseBuilder {
	builder.httpMessage.body = message

	return builder
}

func (builder *HTTPResponseBuilder) build() *HTTPMessage {
	bodyLength := strconv.Itoa(len(builder.httpMessage.body))
	builder.setHeader("content-length", bodyLength)

	if !containsKey(builder.httpMessage.headers, "content-type") {
		builder.setHeader("content-type", "text/plain")
	}

	return builder.httpMessage
}
