package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

func day3() {
	input := utils.GetInput(getInputFile("day3", false))

	o2Bits := reduceBitList(input, true)
	fmt.Println(o2Bits)
	co2Bits := reduceBitList(input, false)
	fmt.Println(co2Bits)

	o2Rating, _ := strconv.ParseInt(o2Bits, 2, 64)
	co2Rating, _ := strconv.ParseInt(co2Bits, 2, 64)
	fmt.Println(o2Rating, co2Rating, o2Rating*co2Rating)
}

func reduceBitList(list []string, useMostCommon bool) string {
	currentBitIndex := 0

	remainingLines := make([]string, len(list))
	copy(remainingLines, list)

	for len(remainingLines) > 1 {
		onesCount := 0

		// get most common bit
		for _, line := range remainingLines {
			if line[currentBitIndex] == 49 {
				onesCount++
			}
		}
		//fmt.Println("ones:", onesCount)
		isOneMostCommon := (onesCount * 2) > len(remainingLines)
		isEqual := (onesCount * 2) == len(remainingLines)

		// trim list to most/least common bit
		var trimmedLines []string
		for _, line := range remainingLines {
			//fmt.Println(line[currentBitIndex])
			switch line[currentBitIndex] {
			case 49:
				// add 1 if 1 most common
				if !isEqual && isOneMostCommon == useMostCommon {
					trimmedLines = append(trimmedLines, line)
				}
				// add 1 if equal and we're looking for most common
				if isEqual && useMostCommon {
					trimmedLines = append(trimmedLines, line)
				}
			case 48:
				// add 0 if 1 not most common
				if !isEqual && isOneMostCommon != useMostCommon {
					trimmedLines = append(trimmedLines, line)
				}
				// add 0 if equal, and we're not looking for most common
				if isEqual && !useMostCommon {
					trimmedLines = append(trimmedLines, line)
				}
			}
		}
		//fmt.Println("list length:", len(list))
		//fmt.Println("remaining lines:", len(remainingLines))
		//fmt.Println("trimmed lines:", len(trimmedLines))
		remainingLines = trimmedLines
		//fmt.Println("new remaining lines:", len(remainingLines))

		//increment bit index
		currentBitIndex++
		if currentBitIndex >= len(remainingLines[0]) {
			currentBitIndex = 0
		}
	}
	//fmt.Println(remainingLines)
	return remainingLines[0]
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
