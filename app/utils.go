package main

import (
	"bufio"
	"fmt"
	"strings"
)

func warn(message string) {
	fmt.Println("[WARN] " + message)
}

func containsKey[K comparable, V any](m map[K]V, k K) bool {
	_, exists := m[k]
	return exists
}

func writeLine(writer *bufio.Writer, message string) error {
	_, err := writer.WriteString(message + "\r\n")
	if err != nil {
		return err
	}

	return nil
}

func readLine(reader *bufio.Reader) string {
	res, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.TrimRight(res, "\r\n")

}
