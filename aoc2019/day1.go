package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

func getFuel(mass float64) (fuel float64) {
	fuel = math.Floor(mass/3) - 2
	if fuel < 0 {
		fuel = 0
	}
	return
}

func getExtraFuel(mass float64) (fuel float64) {
	xFuel := getFuel(mass)
	if xFuel > 0 {
		xFuel += getExtraFuel(xFuel)
	}
	return xFuel
}

func day1() {
	fmt.Println("day 1")
	in := utils.GetInput("aoc2019/day1.txt")

	var totalFuel float64
	for _, v := range in {
		mass, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalln("Input error:", err)
		}
		fuel := getFuel(mass)
		fmt.Println("fuel mass", fuel)
		xFuel := getExtraFuel(fuel)
		fmt.Println("extra mass", xFuel)
		totalFuel += fuel
		totalFuel += xFuel
	}
	fmt.Println(int(totalFuel))
}
