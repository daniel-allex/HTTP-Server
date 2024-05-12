package main

import "strings"

func (builder *HTTPResponseBuilder) compress(request *HTTPMessage) *HTTPResponseBuilder {
	if val, ok := request.headers["accept-encoding"]; ok {
		compressTypes := strings.Split(val, ", ")

		for _, compressType := range compressTypes {
			if compressType == "gzip" {
				return builder.setHeader("content-encoding", compressType)
			}
		}
	}

	return builder
}
