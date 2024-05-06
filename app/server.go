package main

import (
	"fmt"
	"net"
	"os"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read and process data from the client
	// ...

	// Write data back to the client
	// ...

	message := "HTTP/1.1 200 OK\r\n\r\n"
	conn.Write([]byte(message))
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleClient(conn)
}
