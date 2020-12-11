package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
)

func day5() {
	input := utils.GetInput("aoc2020/day5.txt")

	var seatIdList []int
	highestSeatId := 0
	for _, line := range input {
		x := 0
		y := 0
		for i, charInt := range line {
			char := string(charInt)
			if i < 7 {
				switch char {
				case "B":
					x += 1 << (6 - i)
				}
			} else {
				switch char {
				case "R":
					y += 1 << (9 - i)
				}
			}
		}
		seatId := x*8 + y
		seatIdList = append(seatIdList, seatId)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	fmt.Println("Highest ID:", highestSeatId)
	sort.Ints(seatIdList)

	currentSeatId := 0
	for _, sid := range seatIdList {
		if currentSeatId+1 != sid {
			fmt.Println("Missing id before", sid)
		}
		currentSeatId = sid
	}
}
