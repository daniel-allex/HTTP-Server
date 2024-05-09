package main

import "strconv"

type HTTPResponseBuilder struct {
	httpMessage HTTPMessage
}

func createHTTPResponseBuilder() *HTTPResponseBuilder {
	return &HTTPResponseBuilder{createEmptyHTTPMessage(emptyHTTPResponseLine())}
}

func (builder *HTTPResponseBuilder) setVersion(version string) *HTTPResponseBuilder {
	getHTTPResponseLine(builder.httpMessage).version = version

	return builder
}

func (builder *HTTPResponseBuilder) setStatusCode(statusCode int) *HTTPResponseBuilder {
	getHTTPResponseLine(builder.httpMessage).statusCode = statusCode

	return builder
}

func (builder *HTTPResponseBuilder) setStatusText(statusText string) *HTTPResponseBuilder {
	getHTTPResponseLine(builder.httpMessage).statusText = statusText

	return builder
}

func (builder *HTTPResponseBuilder) setHeader(key string, val string) *HTTPResponseBuilder {
	builder.httpMessage.headers[key] = val

	return builder
}

func (builder *HTTPResponseBuilder) setBody(message string) *HTTPResponseBuilder {
	builder.httpMessage.body = message

	return builder
}

func (builder *HTTPResponseBuilder) build() HTTPMessage {
	bodyLength := strconv.Itoa(len(builder.httpMessage.body))
	builder.setHeader("Content-Length", bodyLength)

	if !containsKey(builder.httpMessage.headers, "Content-Type") {
		builder.setHeader("Content-Type", "text/plain")
	}

	return builder.httpMessage
}
