package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func day1() {
	input := utils.GetInput("aoc2021/day1.txt")
	intSlice, err := utils.StringSliceToInts(input)
	if err != nil {
		log.Fatalln("failed to convert input to ints", err)
	}
	fmt.Println(intSlice)

	lastDepth := 0
	totalDeeper := 0

	for i, _ := range intSlice {
		if i < 2 {
			continue
		}
		spreadDepth := intSlice[i-2] + intSlice[i-1] + intSlice[i]
		if lastDepth != 0 && spreadDepth > lastDepth {
			totalDeeper++
		}
		lastDepth = spreadDepth
	}
	fmt.Println(totalDeeper)
}
