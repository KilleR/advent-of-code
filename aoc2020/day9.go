package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func day9() {
	input, err := utils.StringSliceToInts(utils.GetInput("aoc2020/day9.txt"))
	if err != nil {
		log.Fatalln("input failure", err)
	}

	preambleLength := 25

	invalidNum := 0

iterator:
	for loc, line := range input {
		if loc < preambleLength {
			continue
		}

		chunk := input[loc-preambleLength : loc]

		for i := 0; i < preambleLength; i++ {
			for j := i + 1; j < preambleLength; j++ {
				if chunk[i]+chunk[j] == line {
					continue iterator
				}
			}
		}
		fmt.Println(line, "is INVALID")
		invalidNum = line
		break iterator
	}

	//forwards first
contiguity:
	for i := 0; i < len(input); i++ {
		iVal := input[i]
		currentSum := iVal
		currentMin := iVal
		currentMax := iVal
		for j := i + 1; j < len(input); j++ {
			jVal := input[j]
			if jVal < currentMin {
				currentMin = jVal
			}
			if jVal > currentMax {
				currentMax = jVal
			}

			currentSum += jVal

			if currentSum == invalidNum {
				fmt.Println("MATCH!")
				fmt.Println(currentMin+currentMax, "is the crypto vulnerability")
				break contiguity
			}
			if currentSum > invalidNum {
				continue contiguity
			}
		}
	}
}
