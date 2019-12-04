package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handleOpCode(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output OpCodeResult
	}{
		{"halt", []int{99, 2, 3, 4}, OpCodeResult{0, 0, 0, false}},
		{"add", []int{1, 2, 3, 4}, OpCodeResult{7, 4, 4, true}},
		{"mult", []int{2, 4, 6, 8, 10, 12, 14}, OpCodeResult{140, 8, 4, true}},
	}

	a := assert.New(t)
	for _, test := range tests {
		res, _ := handleOpCode(test.input, 0)
		a.Equal(test.output, res, test.name)
	}
}

func Test_intCodeMachine(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"add", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"multiply", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"extended", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"double", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	a := assert.New(t)
	for _, test := range tests {
		a.Equal(test.output, intCodeMachine(test.input), test.name)
	}
}
