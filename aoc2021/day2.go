package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func day2() {
	input := utils.GetInput(getInputFile("day2", false))

	x := 0
	y := 0
	aim := 0

	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			x += magnitude
			y += magnitude * aim
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}

	fmt.Println(x, y)
	fmt.Println(x * y)
}
