package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Rule for the fields
type Rule struct {
	Ranges []Range
	Name   string
}

// Range of a rule
type Range struct {
	Min, Max int
}

func parse(data []string) ([]Rule, []int, [][]int) {
	var rules []Rule
	var ticket []int
	var otherTickets [][]int

	for i := 0; i < len(data); i++ {
		line := data[i]

		if i == 0 {
			for line != "" {
				lineSplit := strings.Split(line, ": ")
				name := lineSplit[0]
				rangesSplit := strings.Split(lineSplit[1], " or ")

				var ranges []Range
				for y := 0; y < len(rangesSplit); y++ {
					rangeSplit := strings.Split(rangesSplit[y], "-")

					min, err := strconv.Atoi(rangeSplit[0])
					check(err)

					max, err := strconv.Atoi(rangeSplit[1])
					check(err)
					ranges = append(ranges, Range{Min: min, Max: max})

				}

				rules = append(rules, Rule{Name: name, Ranges: ranges})

				i++
				line = data[i]
			}
		}

		if strings.HasPrefix(line, "your ticket:") {
			i++
			line = data[i]
			lineSplit := strings.Split(line, ",")
			for y := 0; y < len(lineSplit); y++ {
				val, err := strconv.Atoi(lineSplit[y])
				check(err)

				ticket = append(ticket, val)
			}

			i++
		}

		if strings.HasPrefix(line, "nearby tickets:") {
			i++
			line = data[i]

			for line != "" {
				var ticket []int

				lineSplit := strings.Split(line, ",")
				for y := 0; y < len(lineSplit); y++ {
					val, err := strconv.Atoi(lineSplit[y])
					check(err)

					ticket = append(ticket, val)
				}

				otherTickets = append(otherTickets, ticket)

				i++
				line = data[i]
			}
		}
	}

	return rules, ticket, otherTickets
}

func validateTickets(tickets [][]int, rules []Rule) ([][]int, [][]int, int) {
	var valid [][]int
	var invalid [][]int
	errorRate := 0

	for i := 0; i < len(tickets); i++ {
		ticket := tickets[i]

		isValid := false
		for j := 0; j < len(ticket); j++ {
			value := ticket[j]

			isValid = validateValue(value, rules)
			if !isValid {
				errorRate += value
				break
			}
		}

		if isValid {
			valid = append(valid, ticket)
		} else {
			invalid = append(invalid, ticket)
		}
	}

	return valid, invalid, errorRate
}

func validateValue(value int, rules []Rule) bool {
	for i := 0; i < len(rules); i++ {
		ranges := rules[i].Ranges

		for j := 0; j < len(ranges); j++ {
			rng := ranges[j]

			if value >= rng.Min && value <= rng.Max {
				return true
			}
		}
	}
	return false
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	rules, _, otherTickets := parse(data)

	_, _, errorRate := validateTickets(otherTickets, rules)

	fmt.Println("Error rate:", errorRate)
}
