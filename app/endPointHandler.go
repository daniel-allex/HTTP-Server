package main

import (
	"os"
	"strings"
)

type EndPointData struct {
	endPoint string
	relPath  string
}

func getEndPointData(absPath string) *EndPointData {
	if len(absPath) == 0 || absPath[0] != '/' {
		return nil
	}

	endPoint := ""
	relPath := ""
	if len(absPath) > 1 {
		endPoint, relPath, _ = strings.Cut(absPath[1:], "/")
	}

	return &EndPointData{endPoint: endPoint, relPath: relPath}
}

func rootEndPoint() *HTTPMessage {
	return createSuccessHTTPResponseBuilder().build()
}

func echoEndPoint(relPath string) *HTTPMessage {
	return createSuccessHTTPResponseBuilder().setBody(relPath).build()
}

func userAgentEndPoint(request *HTTPMessage) *HTTPMessage {
	res, ok := request.headers["user-agent"]
	exceptIfNotOk("User agent not in headers", ok)
	return createSuccessHTTPResponseBuilder().setBody(res).build()
}

func filesEndPoint(relPath string, absPath string) *HTTPMessage {
	filePath := absPath + relPath
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return createSuccessHTTPResponseBuilder().
		setBody(string(fileContent)).
		setHeader("Content-Type", "application/octet-stream").
		build()
}

func endPointDataToResponse(endPointData EndPointData, request *HTTPMessage, filePath string) *HTTPMessage {
	if endPointData.endPoint == "echo" {
		return echoEndPoint(endPointData.relPath)
	} else if endPointData.endPoint == "user-agent" {
		return userAgentEndPoint(request)
	} else if endPointData.endPoint == "files" {
		return filesEndPoint(endPointData.relPath, filePath)
	} else if endPointData.endPoint == "" && endPointData.relPath == "" {
		return rootEndPoint()
	} else {
		return createNotFoundHTTPResponseBuilder().build()
	}
}
