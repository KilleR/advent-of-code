package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

const (
	DIR_E = iota
	DIR_S
	DIR_W
	DIR_N
)

func day12() {
	rex := regexp.MustCompile(`([A-Z])([0-9]+)`)

	input := utils.GetInput("aoc2020/day12.txt")

	nPos := 0
	ePos := 0
	heading := DIR_E

	for _, line := range input {
		// code here
		fmt.Println(line)
		matches := rex.FindStringSubmatch(line)
		fmt.Println(matches)

		dir := matches[1]
		magnitude, _ := strconv.Atoi(matches[2])

		switch dir {
		case "E":
			ePos += magnitude
		case "W":
			ePos -= magnitude
		case "N":
			nPos += magnitude
		case "S":
			nPos -= magnitude
		case "R":
			heading = (heading + magnitude/90) % 4
		case "L":
			heading = (heading - magnitude/90) % 4
			if heading < 0 {
				heading += 4
			}
		case "F":
			switch heading {
			case DIR_E:
				ePos += magnitude
			case DIR_W:
				ePos -= magnitude
			case DIR_N:
				nPos += magnitude
			case DIR_S:
				nPos -= magnitude
			}
		}
		fmt.Println("we are at", ePos, nPos, heading)
	}

	fmt.Println("we are at", ePos, nPos)
	eAbs := ePos
	if eAbs < 0 {
		eAbs *= -1
	}
	nAbs := nPos
	if nAbs < 0 {
		nAbs *= -1
	}
	fmt.Println("Manhattan coord", nAbs+eAbs)
}
