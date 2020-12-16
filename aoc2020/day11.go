package main

import (
	"advent-of-code/utils"
	"fmt"
)

type seat string

const (
	SEAT_EMPTY = "L"
	SEAT_FULL  = "#"
	SEAT_NON   = "."
)

func day11() {
	input := utils.GetInput("aoc2020/day11.txt")

	seats := buildSeatMap(input)

	//fmt.Println(seats)

	switched := false
	for {
		//fmt.Println("switch")
		seats, switched = switchSeats(seats)
		if !switched {
			break
		}

		//for _, sr := range seats {
		//	fmt.Println(sr)
		//}
	}

	occupied := 0
	for _, seatRow := range seats {
		for _, seat := range seatRow {
			if seat == SEAT_FULL {
				occupied++
			}
		}
	}

	fmt.Println(occupied, "seats occupied")
}

func buildSeatMap(lines []string) (seats [][]string) {
	for _, line := range lines {
		var seatRow []string
		for _, entry := range line {
			seatRow = append(seatRow, string(entry))
		}
		seats = append(seats, seatRow)
	}

	return
}

func switchSeats(seatMap [][]string) (newSeatMap [][]string, switched bool) {
	var switchMap [][]int
	for x, seatRow := range seatMap {
		var switchRow []int
		for y, _ := range seatRow {
			seatCount := 0
			for i := -1; i <= 1; i++ {
			innerSeatLoop:
				for j := -1; j <= 1; j++ {
					xcoord := x + i
					ycoord := y + j

					if xcoord == x && ycoord == y {
						continue innerSeatLoop
					}

					var adjSeat string

					for {
						if xcoord < 0 || xcoord >= len(seatMap) {
							continue innerSeatLoop
						}
						if ycoord < 0 || ycoord >= len(seatRow) {
							continue innerSeatLoop
						}
						adjSeat = seatMap[xcoord][ycoord]

						if adjSeat != SEAT_NON {
							break
						}
						xcoord += i
						ycoord += j
					}
					if adjSeat == "#" {
						seatCount++
					}

				}
			}
			switchRow = append(switchRow, seatCount)
		}
		switchMap = append(switchMap, switchRow)
	}

	for i, seatRow := range seatMap {
		var newSeatRow []string
		for j, seat := range seatRow {
			switchCount := switchMap[i][j]
			if switchCount == 0 && seat == SEAT_EMPTY {
				newSeatRow = append(newSeatRow, SEAT_FULL)
				switched = true
			} else if switchCount >= 5 && seat == SEAT_FULL {
				newSeatRow = append(newSeatRow, SEAT_EMPTY)
				switched = true
			} else {
				newSeatRow = append(newSeatRow, seat)
			}
		}
		newSeatMap = append(newSeatMap, newSeatRow)
	}

	return
}
