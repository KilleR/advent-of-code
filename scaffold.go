package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var year = time.Now().Year()

func scaffold(day int) {
	goTemplate := fmt.Sprintf(`package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	input := utils.GetInput("aoc%[2]d/day%[1]d/day%[1]d_ex.txt")
  
	for _, line := range input {
    	// code here
		fmt.Println(line)
	}
}
`, day, year)

	err := os.WriteFile(fmt.Sprintf("aoc%d/day%d/day%d.go", year, day, day), []byte(goTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create go file", err)
	}

	exTxtTemplate := fmt.Sprint(`// example text here`)
	err = os.WriteFile(fmt.Sprintf("aoc%d/day%d/day%d_ex.txt", year, day, day), []byte(exTxtTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create example text file", err)
	}

	txtTemplate := fmt.Sprint(`// input text here`)
	err = os.WriteFile(fmt.Sprintf("aoc%d/day%d/day%d.txt", year, day, day), []byte(txtTemplate), 0644)
	if err != nil {
		log.Fatalln("failed to create input text file", err)
	}
}

func main() {
	scaffold(1)
}
