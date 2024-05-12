package main

import (
	"flag"
	"fmt"
)

func responseFromRequest(request *HTTPMessage, filePath string) *HTTPMessage {
	requestLine := getHTTPRequestLine(request)
	endPointData := getEndPointData(requestLine.path)

	if endPointData == nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return endPointDataToResponse(*endPointData, request, filePath)
}

func handleClient(connection *HTTPConnection, filePath string) {
	defer connection.Close()

	for {
		messageContent := connection.nextRequest()
		response := responseFromRequest(messageContent, filePath)
		connection.sendResponse(response)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	filePath := flag.String("directory", "/", "an absolute path to the file system")
	flag.Parse()

	listener := listenTCPConnection("0.0.0.0:4221")
	for {
		conn := acceptHTTPConnection(listener)
		go handleClient(conn, *filePath)
	}
}
