package main

type HTTPResponseBuilder struct {
	httpMessage HTTPMessage
}

func createHTTPResponseBuilder() HTTPResponseBuilder {
	return HTTPResponseBuilder{createEmptyHTTPMessage(emptyHTTPResponseLine())}
}

func getHTTPResponseLine(httpMessage HTTPMessage) *HTTPResponseLine {
	responseLine, ok := httpMessage.startLine.(HTTPResponseLine)
	exceptIfNotOk("Failed to cast HTTPStartLine to HTTPResponseLine", ok)

	return &responseLine
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
