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

func bodyFromMessage(message *HTTPMessage) *string {
	requestLine := getHTTPRequestLine(message)

	body := ""
	if path := pathFromEndPoint("/echo", requestLine.path); path != nil {
		body = strings.TrimLeft(*path, "/")
	} else if pathFromEndPoint("/user-agent", requestLine.path) != nil {
		body = message.headers["user-agent"]
	} else if *pathFromEndPoint("/", requestLine.path) == "" {
	} else {
		return nil
	}

	return &body
}

func responseFromBody(body *string) *HTTPMessage {
	if body != nil {
		return createHTTPResponseBuilder().
			setVersion("HTTP/1.1").
			setStatusCode(200).
			setStatusText("OK").
			setBody(*body).
			build()
	} else {
		return createHTTPResponseBuilder().
			setVersion("HTTP/1.1").
			setStatusCode(404).
			setStatusText("Not Found").
			build()
	}
}

func handleClient(connection *HTTPConnection) {
	defer connection.Close()

	for {
		messageContent := connection.nextRequest()
		body := bodyFromMessage(messageContent)
		response := responseFromBody(body)
		connection.sendResponse(response)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	listener := listenTCPConnection("0.0.0.0:4221")
	for {
		conn := acceptHTTPConnection(listener)
		go handleClient(conn)
	}
}
