package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"sort"
	"time"
)

type adaptor struct {
	joltage                  int
	validAdaptors            []*adaptor
	getValidAdaptorCountMemo int
	getAdaptorChaimMemo      int
}

func newAdaptor(joltage int) (a *adaptor) {
	a = new(adaptor)
	a.getValidAdaptorCountMemo = -1
	a.getAdaptorChaimMemo = -1
	a.joltage = joltage

	return
}

var currentAdaptorCount int

func (a *adaptor) getValidAdaptorCount() int {
	currentAdaptorCount++
	if a.getValidAdaptorCountMemo != -1 {
		return a.getValidAdaptorCountMemo
	}
	if len(a.validAdaptors) == 0 {
		a.getValidAdaptorCountMemo = 1
		return 1
	}
	total := 0

	for _, adaptor := range a.validAdaptors {
		total += adaptor.getValidAdaptorCount()
	}

	a.getValidAdaptorCountMemo = total
	return total
}

func (a adaptor) getAdaptorChain() int {
	if a.getAdaptorChaimMemo != -1 {
		return a.getAdaptorChaimMemo
	}
	if len(a.validAdaptors) == 0 {
		return 1
	}

	skippedMults := 0
	for _, adaptor := range a.validAdaptors {
		if len(adaptor.validAdaptors) > 1 {
			skippedMults++
		}
	}
	//if skippedMults == 1 && len(a.validAdaptors) == 3 {
	//	return 2
	//}
	//if skippedMults == 1 && len(a.validAdaptors) == 2 {
	//	return 1
	//}
	return len(a.validAdaptors)
}

func day10() {
	input, err := utils.StringSliceToInts(utils.GetInput("aoc2020/day10.txt"))
	if err != nil {
		log.Fatalln("input failure", err)
	}

	sort.Ints(input)

	steps := make(map[int]int)

	adaptors := make(map[int]*adaptor)
	adaptors[0] = newAdaptor(0)

	lastJolts := 0
	for _, jolts := range input {
		adaptors[jolts] = newAdaptor(jolts)

		joltDiff := jolts - lastJolts

		_, ok := steps[joltDiff]
		if !ok {
			steps[joltDiff] = 0
		}

		steps[joltDiff]++

		lastJolts = jolts
	}
	steps[3]++ // last step to device

	fmt.Println(steps)

	go func() {
		for {
			fmt.Println(currentAdaptorCount)
			time.Sleep(time.Second)
		}
	}()

	for jolt, adaptor := range adaptors {
		for j := jolt + 1; j <= jolt+3; j++ {
			compatAdaptor, ok := adaptors[j]
			if ok {
				adaptor.validAdaptors = append(adaptor.validAdaptors, compatAdaptor)
			}
		}
	}

	totalChain := 1
	for _, adaptor := range adaptors {
		totalChain *= adaptor.getAdaptorChain()
	}
	fmt.Println("Adaptor chain:", totalChain)

	vac := adaptors[0].getValidAdaptorCount()
	fmt.Println(vac, "valid adaptor combinations")
}
