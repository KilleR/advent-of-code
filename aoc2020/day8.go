package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type op int8

const (
	ACC op = iota
	JMP
	NOP
)

type instruction struct {
	operation  op
	value      int
	isExecuted bool
}

func day8() {
	input := utils.GetInput("aoc2020/day8.txt")

	var instructions []*instruction

	opRex := regexp.MustCompile(`([a-z]{3}) ([-+][0-9]+)`)

	for _, line := range input {
		matches := opRex.FindStringSubmatch(line)
		//fmt.Println(line, matches)
		opString := matches[1]

		var op op
		switch opString {
		case "acc":
			op = ACC
		case "jmp":
			op = JMP
		case "nop":
			op = NOP
		default:
			log.Fatalln("op was not a valid operator", opString)
		}

		value, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln("failed to convert value", matches[2], err)
		}

		newInstruction := new(instruction)
		newInstruction.operation = op
		newInstruction.value = value
		instructions = append(instructions, newInstruction)
	}

	refactorLoc := 0

refactor:
	for {
		flipOp(instructions[refactorLoc])

		pos := 0
		accumulator := 0

	run:
		for {
			if pos == len(instructions) {
				fmt.Println("Successful exit!")
				fmt.Println("Final Acc:", accumulator)
				break refactor
			}
			ins := instructions[pos]
			if ins.isExecuted {
				break run
			}
			ins.isExecuted = true
			switch ins.operation {
			case NOP:
				pos++
				continue run
			case ACC:
				accumulator += ins.value
				pos++
				continue run
			case JMP:
				pos += ins.value
				continue run
			}
		}

		flipOp(instructions[refactorLoc])
		resetInstructions(instructions)
		refactorLoc++
	}

	fmt.Println("Flip at:", refactorLoc)

}

func flipOp(ins *instruction) {
	switch ins.operation {
	case JMP:
		ins.operation = NOP
	case NOP:
		ins.operation = JMP
	}
}

func resetInstructions(instructions []*instruction) {
	for _, ins := range instructions {
		ins.isExecuted = false
	}
}
