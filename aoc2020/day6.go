package main

import (
	"advent-of-code/utils"
	"fmt"
)

func day6() {
	input := utils.GetInput("aoc2020/day6.txt")

	total := 0
	entries := make(map[int32]int)
	groupCount := 0

	for _, line := range input {
		if line == "" {
			total += getGroupCount(entries, groupCount)
			entries = make(map[int32]int)
			groupCount = 0
		} else {
			groupCount++
			for _, charInt := range line {
				_, ok := entries[charInt]
				if !ok {
					entries[charInt] = 0
				}
				entries[charInt]++
			}
		}
	}
	fmt.Println(total)
}

func getGroupCount(entries map[int32]int, groupCount int) (total int) {
	for _, count := range entries {
		if count == groupCount {
			total++
		}
	}
	return total
}
