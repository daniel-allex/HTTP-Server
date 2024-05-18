package main

import (
	"bytes"
	"compress/gzip"
	"strings"
)

type compressFunc func(string) (string, error)

func gzipCompression(body string) (string, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(body))
	if err != nil {
		return "", err
	}

	err = gz.Close()
	if err != nil {
		return "", err
	}

	res := buf.String()
	return res, nil
}

var compressionFuncs = map[string]compressFunc{
	"gzip": gzipCompression,
}

func (builder *HTTPResponseBuilder) tryCompression(compressType string) (*HTTPResponseBuilder, bool) {
	if compressString, ok := compressionFuncs[compressType]; ok {
		res, err := compressString(builder.httpMessage.body)

		if err != nil {
			warn("Failed to compress with " + compressType + ": " + err.Error())
			return nil, false
		}

		return builder.setHeader("content-encoding", compressType).setBody(res), true
	}

	return builder, false
}

func (builder *HTTPResponseBuilder) compress(request *HTTPMessage) *HTTPResponseBuilder {
	if val, ok := request.headers["accept-encoding"]; ok {
		compressTypes := strings.Split(val, ", ")

		for _, compressType := range compressTypes {
			if _, ok = builder.tryCompression(compressType); ok {
				return builder
			}
		}
	}

	return builder
}
