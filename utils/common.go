package utils

import (
	"fmt"
	"os"
	"strings"
)

func CheckAndExit(err error) {
	if err != nil {
		panic(err)
	}
}

func Exit(message string, code int) {
	if strings.TrimSpace(message) == "" {
		message = "No message"
	}
	fmt.Println(message)
	os.Exit(code)
}
