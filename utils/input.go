package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func GetInput(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}

func fromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}
