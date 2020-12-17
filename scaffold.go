package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var year = 2020

func scaffold(day int) {
	goTemplate := fmt.Sprintf(`package main

import (
	"advent-of-code/utils"
	"fmt"
)

func day%[1]d() {
	input := utils.GetInput("aoc%[2]d/day%[1]d_ex.txt")
  
	for _, line := range input {
    	// code here
		fmt.Println(line)
	}
}
`, day, year)

	err := ioutil.WriteFile(fmt.Sprintf("aoc%d/day%d.go", year, day), []byte(goTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create go file", err)
	}

	exTxtTemplate := fmt.Sprint(`// example text here`)
	err = ioutil.WriteFile(fmt.Sprintf("aoc%d/day%d_ex.txt", year, day), []byte(exTxtTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create example text file", err)
	}

	txtTemplate := fmt.Sprint(`// input text here`)
	err = ioutil.WriteFile(fmt.Sprintf("aoc%d/day%d.txt", year, day), []byte(txtTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create input text file", err)
	}
}

func main() {
	scaffold(12)
}
