package main

import (
	"bytes"
	"compress/gzip"
	"strings"
)

type compressFunc func(string) *string

func gzipCompression(body string) *string {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(body))
	gz.Close()

	if err != nil {
		return nil
	}

	res := buf.String()
	return &res
}

var compressionFuncs = map[string]compressFunc{
	"gzip": gzipCompression,
}

func (builder *HTTPResponseBuilder) tryCompression(compressType string) (*HTTPResponseBuilder, bool) {
	if compressString, ok := compressionFuncs[compressType]; ok {
		res := compressString(builder.httpMessage.body)

		if res != nil {
			return builder.setHeader("content-encoding", compressType).setBody(*res), true
		}
	}

	return builder, false
}

func (builder *HTTPResponseBuilder) compress(request *HTTPMessage) *HTTPResponseBuilder {
	if val, ok := request.headers["accept-encoding"]; ok {
		compressTypes := strings.Split(val, ", ")

		for _, compressType := range compressTypes {
			if _, ok = builder.tryCompression(compressType); ok {
				return builder
			} else {
				warn("Failed to compress with " + compressType)
			}
		}
	}

	return builder
}
