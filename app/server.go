package main

import (
	"fmt"
	"strings"
)

func handleClient(connection *HTTPConnection) {
	defer connection.Close()

	println("got here")
	messageContent := connection.nextRequest()
	println("got request")
	requestLine := getHTTPRequestLine(messageContent)
	_, body, _ := strings.Cut(requestLine.version, "/echo/")

	response := createHTTPResponseBuilder().
		setVersion("HTTP/1.1").
		setStatusCode(200).
		setStatusText("OK").
		setBody(body).
		build()

	connection.sendResponse(response)

	println("got to end")
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	conn := connectHTTP("tcp", "0.0.0.0:4221")

	handleClient(conn)

}
