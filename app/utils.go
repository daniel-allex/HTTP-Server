package main

import (
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
