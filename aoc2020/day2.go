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
	var passwords []string
	type minMax struct {
		min int
		max int
	}
	rules := make(map[string]minMax)

	for _, line := range input {
		matches := pwRex.FindStringSubmatch(line)
		password := matches[4]
		passwords = append(passwords, password)

		min, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalln("couldn't parse min value", matches[1], err)
		}

		max, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln("couldn't parse min value", matches[2], err)
		}

		rules[matches[3]] = minMax{
			min: min,
			max: max,
		}

		pwCharCount := make(map[string]int)
		for _, charInt := range password {
			char := string(charInt)
			_, ok := pwCharCount[char]
			if !ok {
				pwCharCount[char] = 0
			}
			pwCharCount[char]++
		}

		fmt.Println(matches)
		fmt.Println(pwCharCount)
		count, ok := pwCharCount[matches[3]]
		if !ok || count < min || count > max {
			fmt.Println("INVALID")
		} else {
			validCount++
		}

	}

	fmt.Println("Valid Count:", validCount)
}
