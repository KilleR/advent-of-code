package main

import (
	"advent-of-code/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type OpCode int

const (
	HALT OpCode = 99
	ADD  OpCode = 1
	MULT OpCode = 2
)

type OpCodeResult struct {
	result    int
	resultLoc int
	increment int
	cont      bool
}

func handleOpCode(in []int, loc int) (out OpCodeResult, err error) {
	switch in[loc] {
	case 1:
		out.cont = true
		out.result = in[in[loc+1]] + in[in[loc+2]]
		out.increment = 4
		out.resultLoc = in[loc+3]
		return
	case 2:
		out.cont = true
		out.result = in[in[loc+1]] * in[in[loc+2]]
		out.increment = 4
		out.resultLoc = in[loc+3]
		return
	case 99:
		out.cont = false
		return
	default:
		out.cont = false
		err = errors.New("invalid OpCode " + string(in[0]))
		return
	}
}

func intCodeMachine(in []int) (out []int) {
	out = make([]int, len(in))
	copy(out, in)
	for i := 0; i < len(out); {
		oldI := i

		opRes, err := handleOpCode(out, i)
		if err != nil {
			panic("bad OpCode response: " + err.Error())
		}

		if !opRes.cont {
			break
		}

		i += opRes.increment
		out[opRes.resultLoc] = opRes.result

		if i == oldI {
			panic("infinite loop detected")
		}
	}
	return
}

func day2() {
	input := utils.GetInput("aoc2019/day2.txt")
	inStrings := strings.Split(input[0], ",")

	in := make([]int, len(inStrings))
	for i, v := range inStrings {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic("not a string: " + err.Error())
		}
		in[i] = val
	}

	var noun, verb int

loop:
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {

			in[1] = noun
			in[2] = verb

			out := intCodeMachine(in)
			res := out[0]
			if res == 19690720 {
				break loop
			}
		}
	}

	fmt.Println(100*noun + verb)
}
