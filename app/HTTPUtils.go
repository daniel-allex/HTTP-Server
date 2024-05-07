package main

import (
	"bufio"
	"strings"
)

func write(writer bufio.Writer, message string) {
	_, err := writer.WriteString(message)
	validateResult("Failed to write string", err)
}

func writeLine(writer bufio.Writer, message string) {
	write(writer, message+"\r\n")
}

func readLine(reader bufio.Reader) string {
	res, err := reader.ReadString('\n')
	validateResult("Failed to read string", err)

	return strings.Trim(res, "\r")
}
