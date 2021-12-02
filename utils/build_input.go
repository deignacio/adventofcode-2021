package utils

import (
	"os"
	"strings"
)

func ReadInput(inputPath string) []byte {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	return data
}

func AsInputList(rawBytes []byte) []string {
	rawString := string(rawBytes)
	return strings.Split(rawString, "\n")
}
