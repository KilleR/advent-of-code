package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type bag struct {
	bags                 map[string]*bag
	bagsCount            map[string]int
	getContainedBagsMemo int
	canContainBagMemo    map[string]bool
	contentStrings       []string
}

func newBag() (newBag *bag) {
	newBag = new(bag)
	newBag.bags = make(map[string]*bag)
	newBag.bagsCount = make(map[string]int)
	newBag.getContainedBagsMemo = -1
	newBag.canContainBagMemo = make(map[string]bool)
	return
}

func (b bag) getContainedBags() (total int) {
	if b.getContainedBagsMemo != -1 {
		return b.getContainedBagsMemo
	}

	for innerBagColour, innerBag := range b.bags {
		total += b.bagsCount[innerBagColour]
		total += innerBag.getContainedBags() * b.bagsCount[innerBagColour]
	}
	b.getContainedBagsMemo = total
	return
}

func (b bag) canContainBag(colour string) bool {
	ccb, ok := b.canContainBagMemo[colour]
	if ok {
		return ccb
	}

	_, ok = b.bags[colour]
	if ok {
		b.canContainBagMemo[colour] = true
		return true
	}
	for _, innerBag := range b.bags {
		if innerBag.canContainBag(colour) {
			b.canContainBagMemo[colour] = true
			return true
		}
	}

	b.canContainBagMemo[colour] = false
	return false
}

func day7() {
	input := utils.GetInput("aoc2020/day7.txt")

	bagRex := regexp.MustCompile(`([^ ]+ [^ ]+) bags? contain (.+)`)
	contentRex := regexp.MustCompile(`([0-9]+) ([^ ]+ [^ ]+) bags?`)

	allBags := make(map[string]*bag)

	// mamke the bags
	for _, line := range input {
		matches := bagRex.FindStringSubmatch(line)
		//fmt.Println("line:", line, matches)
		bagColour := matches[1]
		//fmt.Println("Bag:", bagColour)
		bagContents := matches[2]
		//fmt.Println("Contents:", bagContents)

		contentParts := strings.Split(bagContents, ",")

		thisBag := newBag()
		thisBag.contentStrings = contentParts
		allBags[bagColour] = thisBag
	}

	// stitch the bags together
	for bagColour, bag := range allBags {
		for _, content := range bag.contentStrings {
			matches := contentRex.FindStringSubmatch(content)
			//fmt.Println(bagColour, matches)
			if len(matches) == 0 {
				continue
			}
			innerBagCount, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalln("failed to convert bag count to number", matches[1], err)
			}
			innerBagColour := matches[2]
			innerBag, ok := allBags[innerBagColour]
			if !ok {
				log.Fatalln("no main bag for ", bagColour, innerBagColour)
			}

			bag.bags[innerBagColour] = innerBag
			bag.bagsCount[innerBagColour] = innerBagCount
		}
	}

	canContainShinyGold := 0
	for _, bag := range allBags {
		//fmt.Println(bagColour, bag.bags, bag.bagsCount)
		//fmt.Println(bag.getContainedBags())
		if bag.canContainBag("shiny gold") {
			canContainShinyGold++
			//fmt.Println("SHINY GOLD!")
		}
	}
	fmt.Println(canContainShinyGold, "can contain shiny gold")

	startTwo := time.Now()
	fmt.Println(allBags["shiny gold"].getContainedBags(), "bags in shiny gold")
	fmt.Println(time.Since(startTwo))
}
