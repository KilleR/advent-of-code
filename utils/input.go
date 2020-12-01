package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
		fmt.Println(v)
		currentInt, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		out[i] = currentInt
	}
	return
}
