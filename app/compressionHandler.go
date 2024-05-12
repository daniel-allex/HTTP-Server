package main

func (builder *HTTPResponseBuilder) compress(request *HTTPMessage) *HTTPResponseBuilder {
	if compressType, ok := request.headers["accept-encoding"]; ok {
		if compressType == "gzip" {
			return builder.setHeader("content-encoding", compressType)
		}
	}

	return builder
}
