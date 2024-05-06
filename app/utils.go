package main

import (
	"fmt"
	"os"
)

func validateResult(message string, err error) {
	if err != nil {
		fmt.Println(message + ": " + err.Error())
		os.Exit(1)
	}
}
