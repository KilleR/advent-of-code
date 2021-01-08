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

	sPosN := 0
	sPosE := 0
	wpPosN := 1
	wpPosE := 10

	for _, line := range input {
		// code here
		//fmt.Println(line)
		matches := rex.FindStringSubmatch(line)
		//fmt.Println(matches)

		dir := matches[1]
		magnitude, _ := strconv.Atoi(matches[2])

		switch dir {
		case "E":
			wpPosE += magnitude
		case "W":
			wpPosE -= magnitude
		case "N":
			wpPosN += magnitude
		case "S":
			wpPosN -= magnitude
		case "R":
			rotCount := (magnitude / 90) % 4
			for i := 0; i < rotCount; i++ {
				oldE := wpPosE
				oldN := wpPosN
				wpPosE = oldN
				wpPosN = -oldE
			}
		case "L":
			rotCount := (magnitude / 90) % 4
			for i := 0; i < rotCount; i++ {
				oldE := wpPosE
				oldN := wpPosN
				wpPosE = -oldN
				wpPosN = oldE
			}
		case "F":
			sPosN += wpPosN * magnitude
			sPosE += wpPosE * magnitude
		}
		//fmt.Println("waypoint at", wpPosE, wpPosN)
		//fmt.Println("ship at", sPosE, sPosN)
	}

	fmt.Println("waypoint at", wpPosE, wpPosN)
	fmt.Println("ship at", sPosE, sPosN)
	eAbs := sPosE
	if eAbs < 0 {
		eAbs *= -1
	}
	nAbs := sPosN
	if nAbs < 0 {
		nAbs *= -1
	}
	fmt.Println("Manhattan coord", nAbs+eAbs)
}
