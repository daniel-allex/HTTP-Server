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

func filesPostEndPoint(filePath string, request *HTTPMessage) *HTTPMessage {
	err := os.WriteFile(filePath, []byte(request.body), 0666)

	if err != nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return createSuccessHTTPResponseBuilder().
		setStatusCode(201).
		setStatusText("Created").
		build()
}

func filesGetEndPoint(filePath string, request *HTTPMessage) *HTTPMessage {
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return createSuccessHTTPResponseBuilder().
		setBody(string(fileContent)).
		setHeader("Content-Type", "application/octet-stream").
		build()
}

func filesEndPoint(filePath string, request *HTTPMessage) *HTTPMessage {
	requestLine := getHTTPRequestLine(request)

	if requestLine.method == "POST" {
		return filesPostEndPoint(filePath, request)
	} else if requestLine.method == "GET" {
		return filesGetEndPoint(filePath, request)
	}

	return createNotFoundHTTPResponseBuilder().build()
}

func endPointDataToResponse(endPointData EndPointData, request *HTTPMessage, filePath string) *HTTPMessage {
	if endPointData.endPoint == "echo" {
		return echoEndPoint(endPointData.relPath)
	} else if endPointData.endPoint == "user-agent" {
		return userAgentEndPoint(request)
	} else if endPointData.endPoint == "files" {
		return filesEndPoint(filePath+"/"+endPointData.relPath, request)
	} else if endPointData.endPoint == "" && endPointData.relPath == "" {
		return rootEndPoint()
	} else {
		return createNotFoundHTTPResponseBuilder().build()
	}
}
