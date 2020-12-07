package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func day2() {
	pwRex := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

	input := utils.GetInput("aoc2020/day2.txt")

	validCount := 0

	for _, line := range input {
		matches := pwRex.FindStringSubmatch(line)
		password := matches[4]
		char := matches[3]

		min, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalln("couldn't parse min value", matches[1], err)
		}

		max, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln("couldn't parse min value", matches[2], err)
		}

		fmt.Println(password, char, min, max)

		firstChar := string(password[min-1])
		secondChar := string(password[max-1])
		if firstChar != secondChar && (firstChar == char || secondChar == char) {
			validCount++
			fmt.Println("VALID")
		}

	}

	fmt.Println("Valid Count:", validCount)
}
