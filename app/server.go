package main

import (
	"fmt"
	"strings"
)

func handleClient(connection *HTTPConnection) {
	defer connection.Close()

	messageContent := connection.nextRequest()
	requestLine := getHTTPRequestLine(messageContent)

	body := ""
	success := false
	if requestLine.path == "/" {
		body = ""
		success = true
	} else {
		_, body, success = strings.Cut(requestLine.path, "/echo/")
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
