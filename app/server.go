package main

import (
	"flag"
	"fmt"
	"net"
)

func responseFromRequest(request *HTTPMessage, filePath string) *HTTPMessage {
	requestLine := getHTTPRequestLine(request)
	endPointData := getEndPointData(requestLine.path)

	if endPointData == nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return endPointDataToResponseBuilder(*endPointData, request, filePath).compress(request).build()
}

func handleClient(connection *HTTPConnection, filePath string) {
	defer connection.Close()

	request := connection.nextRequest()
	response := responseFromRequest(request, filePath)
	connection.sendResponse(response)
}

func closeListener(listener *net.Listener) {
	err := (*listener).Close()
	validateResult("Failed to close listener", err)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	filePath := flag.String("directory", "/", "an absolute path to the file system")
	flag.Parse()

	listener := listenTCPConnection("0.0.0.0:4221")
	defer closeListener(listener)

	for {
		conn := acceptHTTPConnection(listener)
		go handleClient(conn, *filePath)
	}
}
