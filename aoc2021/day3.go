package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

func day3() {
	input := utils.GetInput(getInputFile("day3", true))

	bitCount := 0
	inputLength := 0
	currentBitIndex := 0
	for _, line := range input {
		inputLength++

		if line[currentBitIndex] == 49 {
			bitCount++
		}
	}
	fmt.Println(bitCount)
}

func day3_part1() {
	input := utils.GetInput(getInputFile("day3", false))

	bitCount := make([]int, len(input[0]))
	inputLength := 0
	for _, line := range input {
		inputLength++

		for charIndex, char := range line {
			if char == 49 {
				bitCount[charIndex]++
			}
		}
	}
	fmt.Println(bitCount)

	var gammaRate, epsillonRate string
	for _, count := range bitCount {
		if count*2 > inputLength {
			gammaRate += "1"
			epsillonRate += "0"
		} else {
			gammaRate += "0"
			epsillonRate += "1"
		}
	}

	fmt.Println(gammaRate, epsillonRate)
	resolvedGammaRate, _ := strconv.ParseInt(gammaRate, 2, 64)
	resolvedEpsillonRate, _ := strconv.ParseInt(epsillonRate, 2, 64)

	fmt.Println("gamma rate:", resolvedGammaRate)
	fmt.Println("epsillon rate:", resolvedEpsillonRate)
	fmt.Println("power consumption:", resolvedGammaRate*resolvedEpsillonRate)
}
