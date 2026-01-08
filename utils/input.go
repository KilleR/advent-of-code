package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetInputFile(year string, day string, useExampleFile bool) (inputString string) {
	inputString = "aoc" + year + "/" + day
	if useExampleFile {
		inputString += "_ex"
	}
	inputString += ".txt"

	return inputString
}

func GetInput(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\n")
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

func StringSliceToInts(in []string) (out []int, err error) {
	out = make([]int, len(in))
	var currentInt int
	for i, v := range in {
		currentInt, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		out[i] = currentInt
	}
	return
}

func StringSliceToIntsWithPrefix(in []string, prefixLen int) (out []int, outPrefix []string, err error) {
	out = make([]int, len(in))
	outPrefix = make([]string, len(in))
	var currentInt int
	for i, v := range in {
		currentPrefix := v[:prefixLen]
		currentInt, err = strconv.Atoi(v[prefixLen:])
		if err != nil {
			return
		}
		out[i] = currentInt
		outPrefix[i] = currentPrefix
	}
	return
}
