package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func day1() {
	input := utils.GetInput("aoc2025/day1.txt")
	intSlice, prefixSlice, err := utils.StringSliceToIntsWithPrefix(input, 1)
	if err != nil {
		log.Fatalln("failed to convert input to ints", err)
	}
	//fmt.Println(intSlice)
	//fmt.Println(prefixSlice)

	dial := 50
	zeroCount := 0

	for i, _ := range intSlice {
		if dial == 0 && prefixSlice[i] == "R" && prefixSlice[i-1] == "L" {
			zeroCount++
			//fmt.Printf("At 0 before %s %d\n", prefixSlice[i], intSlice[i])
		}
		if dial == 0 && prefixSlice[i] == "L" && prefixSlice[i-1] == "R" {
			zeroCount--
			//fmt.Printf("Compensate for start at 0 - %s %d\n", prefixSlice[i], intSlice[i])
		}
		switch prefixSlice[i] {
		case "L":
			dial -= intSlice[i]
			//fmt.Printf("Dial Left %d to %d\n", intSlice[i], dial)
		case "R":
			dial += intSlice[i]
			//fmt.Printf("Dial Right %d to %d\n", intSlice[i], dial)
		default:
			log.Fatalln("unknown prefix type", prefixSlice[i])
		}
		for dial > 99 {
			dial -= 100
			zeroCount++
			//fmt.Printf("Hit 0 during %s %d\n", prefixSlice[i], intSlice[i])
		}
		for dial < 0 {
			dial += 100
			zeroCount++
			//fmt.Printf("Hit 0 during %s %d\n", prefixSlice[i], intSlice[i])
		}
	}
	fmt.Println(dial)
	fmt.Println(prefixSlice[len(prefixSlice)-1])
	fmt.Println(zeroCount)
}
