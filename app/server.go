package main

import (
	"fmt"
)

func handleClient(connection HTTPConnection) {
	defer connection.Close()

	messageContent := connection.nextRequest()

	if messageContent.path == "/" {
		connection.WriteSuccessResponse("OK")
	} else {
		connection.WriteFailedResponse("Not Found")
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	conn := connectHTTP("tcp", "0.0.0.0:4221")

	handleClient(conn)

}
