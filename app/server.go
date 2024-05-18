package main

import (
	"flag"
	"fmt"
	"net"
)

func responseFromRequest(request *HTTPMessage, filePath string) *HTTPMessage {
	requestLine, ok := getHTTPRequestLine(request)
	if !ok {
		return createNotFoundHTTPResponseBuilder().build()
	}

	endPointData, err := getEndPointData(requestLine.path)

	if err != nil {
		return createNotFoundHTTPResponseBuilder().build()
	}

	return endPointDataToResponseBuilder(endPointData, request, filePath).compress(request).build()
}

func handleClient(connection *HTTPConnection, filePath string) {
	defer connection.Close()

	request, err := connection.nextRequest()
	if err != nil {
		fmt.Println("failed to respond to request: %w", err)
	}

	response := responseFromRequest(request, filePath)
	err = connection.sendResponse(response)
	if err != nil {
		fmt.Println("failed to respond to request: %w", err)
	}
}

func closeListener(listener *net.Listener) {
	err := (*listener).Close()

	if err != nil {
		fmt.Println("failed to close listener: %w", err)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	filePath := flag.String("directory", "/", "an absolute path to the file system")
	flag.Parse()

	listener, err := listenTCPConnection("0.0.0.0:4221")
	if err != nil {
		fmt.Println("failed to listen for TCP connection: %w", err)
	}

	defer closeListener(listener)

	for {
		conn, err := acceptHTTPConnection(listener)
		if err != nil {
			fmt.Println("failed to accept HTTP connection: %w", err)
		}

		go handleClient(conn, *filePath)
	}
}
