package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"sort"
)

func day1() {
	input := utils.GetInput("aoc2020/day1.txt")
	intSlice, err := utils.StringSliceToInts(input)
	if err != nil {
		log.Fatalln("failed to convert input to ints", err)
	}
	fmt.Println(intSlice)

	pair := getSumFor(2020, intSlice)
	fmt.Println(pair)

	fmt.Println(pair[0] * pair[1] * pair[2])
}

func getSumFor(targetNumber int, sourceSlice []int) (out []int) {
loop:
	for _, i := range sourceSlice {
		for _, j := range sourceSlice {
			for _, k := range sourceSlice {
				if i+j+k == targetNumber {
					out = []int{i, j, k}
					break loop
				}
			}
		}
	}

	sort.Ints(out)
	return
}
