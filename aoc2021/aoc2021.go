package main

import (
	"advent-of-code/utils"
	"fmt"
	"time"
)

const Year = "2021"

func getInputFile(day string, useExampleFile bool) (inputString string) {
	return utils.GetInputFile(Year, day, useExampleFile)
}

func main() {

	fmt.Println("Running ... ")
	start := time.Now()

	day4()

	fmt.Println("done in", time.Since(start))
}
