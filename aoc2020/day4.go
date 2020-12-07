package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func day4() {
	/*
		ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
		byr:1937 iyr:2017 cid:147 hgt:183cm
	*/
	passportMap := make(map[string]*regexp.Regexp)
	passportMap["byr"] = regexp.MustCompile(`byr:([0-9]{4})`)
	passportMap["iyr"] = regexp.MustCompile(`iyr:([0-9]{4})`)
	passportMap["eyr"] = regexp.MustCompile(`eyr:([0-9]{4})`)
	passportMap["hgt"] = regexp.MustCompile(`hgt:([0-9]{2,3})([a-z]{2})`)
	passportMap["hcl"] = regexp.MustCompile(`hcl:(#)([0-9a-f]{6})`)
	passportMap["ecl"] = regexp.MustCompile(`ecl:([a-z]{3})`)
	passportMap["pid"] = regexp.MustCompile(`pid:([0-9]{9})`)
	passportMap["cid"] = regexp.MustCompile(`cid:([0-9]+)`)

	input := utils.GetInput("aoc2020/day4.txt")

	ppMap := make(map[string]string)
	validPassports := 0
	for _, line := range input {
		//fmt.Println(line)

		if line == "" {
			passportIsValid := true
			//fmt.Println("pp map", ppMap)
			for key, _ := range passportMap {
				if key == "cid" {
					continue
				}

				_, ok := ppMap[key]
				if !ok {
					passportIsValid = false
				}
			}
			if passportIsValid {
				fmt.Println(ppMap, "VALID!")
				validPassports++
			}

			ppMap = make(map[string]string)
		}

		for key, rex := range passportMap {
			matches := rex.FindStringSubmatch(line)
			if len(matches) != 0 {
				//fmt.Println("found", key, matches)

				isValidMatch := false
				switch key {
				case "byr":
					year, _ := strconv.Atoi(matches[1])
					if year >= 1920 && year <= 2002 {
						isValidMatch = true
					}
				case "iyr":
					year, _ := strconv.Atoi(matches[1])
					if year >= 2010 && year <= 2020 {
						isValidMatch = true
					}
				case "eyr":
					year, _ := strconv.Atoi(matches[1])
					if year >= 2020 && year <= 2030 {
						isValidMatch = true
					}
				case "hgt":
					hgt, _ := strconv.Atoi(matches[1])
					unit := matches[2]
					if unit == "cm" && hgt >= 150 && hgt <= 193 {
						isValidMatch = true
					} else if unit == "in" && hgt >= 59 && hgt <= 76 {
						isValidMatch = true
					}
				case "hcl":
					isValidMatch = true
				case "ecl":
					validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
					for _, colour := range validEyeColors {
						if colour == matches[1] {
							isValidMatch = true
							break
						}
					}
				case "pid":
					isValidMatch = true
				}

				if isValidMatch {
					ppMap[key] = matches[0]
				}
			}
		}
	}

	passportIsValid := true
	for key, _ := range passportMap {
		if key == "cid" {
			continue
		}
		_, ok := ppMap[key]
		if !ok {
			passportIsValid = false
		}
	}
	if passportIsValid {
		fmt.Println(ppMap, "VALID!")
		validPassports++
	}

	ppMap = make(map[string]string)

	fmt.Println("Valid passports:", validPassports)
}
