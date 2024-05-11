package main

import (
	"fmt"
	"strings"
)

func pathFromEndPoint(endPoint string, route string) *string {
	_, path, success := strings.Cut(route, endPoint)

	if success {
		return &path
	} else {
		return nil
	}
}

func handleClient(connection *HTTPConnection) {
	defer connection.Close()

	messageContent := connection.nextRequest()
	requestLine := getHTTPRequestLine(messageContent)

	body := ""
	success := true
	if path := pathFromEndPoint("/echo", requestLine.path); path != nil {
		body = strings.TrimLeft(*path, "/")
	} else if pathFromEndPoint("/user-agent", requestLine.path) != nil {
		body = messageContent.headers["user-agent"]
	} else if *pathFromEndPoint("/", requestLine.path) == "" {
	} else {
		success = false
	}

	var response *HTTPMessage
	if success {
		response = createHTTPResponseBuilder().
			setVersion("HTTP/1.1").
			setStatusCode(200).
			setStatusText("OK").
			setBody(body).
			build()
	} else {
		response = createHTTPResponseBuilder().
			setVersion("HTTP/1.1").
			setStatusCode(404).
			setStatusText("Not Found").
			build()
	}

	connection.sendResponse(response)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	conn := connectHTTP("tcp", "0.0.0.0:4221")

	handleClient(conn)

}
