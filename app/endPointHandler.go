package main

import (
	"errors"
	"os"
	"strings"
)

type EndPointData struct {
	endPoint string
	relPath  string
}

func getEndPointData(absPath string) (EndPointData, error) {
	if len(absPath) == 0 || absPath[0] != '/' {
		return EndPointData{}, errors.New("invalid end point")
	}

	endPoint := ""
	relPath := ""
	if len(absPath) > 1 {
		endPoint, relPath, _ = strings.Cut(absPath[1:], "/")
	}

	return EndPointData{endPoint: endPoint, relPath: relPath}, nil
}

func rootEndPoint() *HTTPResponseBuilder {
	return createSuccessHTTPResponseBuilder()
}

func echoEndPoint(relPath string) *HTTPResponseBuilder {
	return createSuccessHTTPResponseBuilder().setBody(relPath)
}

func userAgentEndPoint(request *HTTPMessage) *HTTPResponseBuilder {
	res, ok := request.headers["user-agent"]

	if !ok {
		return createNotFoundHTTPResponseBuilder().setBody("user-agent not a valid header")
	}

	return createSuccessHTTPResponseBuilder().setBody(res)
}

func filesPostEndPoint(filePath string, request *HTTPMessage) *HTTPResponseBuilder {
	err := os.WriteFile(filePath, []byte(request.body), 0666)

	if err != nil {
		return createNotFoundHTTPResponseBuilder().setBody("failed to write file")
	}

	return createSuccessHTTPResponseBuilder().
		setStatusCode(201).
		setStatusText("Created")
}

func filesGetEndPoint(filePath string, request *HTTPMessage) *HTTPResponseBuilder {
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		return createNotFoundHTTPResponseBuilder()
	}

	return createSuccessHTTPResponseBuilder().
		setBody(string(fileContent)).
		setHeader("Content-Type", "application/octet-stream")
}

func filesEndPoint(filePath string, request *HTTPMessage) *HTTPResponseBuilder {
	requestLine, ok := getHTTPRequestLine(request)

	if !ok {
		return createNotFoundHTTPResponseBuilder()
	}

	if requestLine.method == "POST" {
		return filesPostEndPoint(filePath, request)
	} else if requestLine.method == "GET" {
		return filesGetEndPoint(filePath, request)
	}

	return createNotFoundHTTPResponseBuilder()
}

func endPointDataToResponseBuilder(endPointData EndPointData, request *HTTPMessage, filePath string) *HTTPResponseBuilder {
	if endPointData.endPoint == "echo" {
		return echoEndPoint(endPointData.relPath)
	} else if endPointData.endPoint == "user-agent" {
		return userAgentEndPoint(request)
	} else if endPointData.endPoint == "files" {
		return filesEndPoint(filePath+"/"+endPointData.relPath, request)
	} else if endPointData.endPoint == "" && endPointData.relPath == "" {
		return rootEndPoint()
	} else {
		return createNotFoundHTTPResponseBuilder()
	}
}
