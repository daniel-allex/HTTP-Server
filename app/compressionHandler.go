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

func (builder *HTTPResponseBuilder) compress(request *HTTPMessage) *HTTPResponseBuilder {
	compressionFuncs := map[string]compressFunc{"gzip": gzipCompression}

	if val, ok := request.headers["accept-encoding"]; ok {
		compressTypes := strings.Split(val, ", ")

		for _, compressType := range compressTypes {
			if compressString, ok := compressionFuncs[compressType]; ok {
				res := compressString(builder.httpMessage.body)

				if res != nil {
					return builder.setHeader("content-encoding", compressType).setBody(*res)
				} else {
					warn("Failed to compress with " + compressType + "... Trying other options")
				}
			} else {
				warn(compressType + " not available")
			}
		}
	}

	return builder
}
