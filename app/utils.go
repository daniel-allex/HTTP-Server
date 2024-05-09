package main

import (
	"bufio"
	"fmt"
	"os"
)

func validateResult(message string, err error) {
	if err != nil {
		throwException(message + ": " + err.Error())
	}
}

func exceptIfNotOk(message string, ok bool) {
	if !ok {
		throwException(message)
	}
}

func throwException(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func containsKey[K comparable, V any](m map[K]V, k K) bool {
	_, exists := m[k]
	return exists
}

func write(writer *bufio.Writer, message string) {
	_, err := writer.WriteString(message)
	validateResult("Failed to write string", err)
}

func writeLine(writer *bufio.Writer, message string) {
	write(writer, message+"\r\n")
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
