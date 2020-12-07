package main

import (
	"advent-of-code/utils"
	"fmt"
)

func day3() {
	input := utils.GetInput("aoc2020/day3.txt")

	paths := []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var treeCounts []int
	for _, path := range paths {
		loc := 0
		treeCount := 0
		for i, line := range input {

			if i%path.down == 0 {
				if string(line[loc%len(line)]) == "#" {
					treeCount++
				}

				loc += path.right
			}
		}

		treeCounts = append(treeCounts, treeCount)
	}

	fmt.Println(treeCounts)

	tot := 1
	for _, tc := range treeCounts {
		tot *= tc
	}
	fmt.Println(tot)
}
