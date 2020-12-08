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

func findPosition(instruction string, data []string) map[int]bool {
	var positions = make(map[int]bool)
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], instruction) {
			positions[i] = true
		}
	}
	return positions
}

func runSwap(from string, to string, data []string, positions map[int]bool) (bool, int) {
	accumulator := 0
	exited := false

	for postion := range positions {
		if exited {
			break
		}
		accumulator = 0
		var vistedInstructions = make(map[int]bool)
		for i := 0; i < len(data); i++ {
			if _, visted := vistedInstructions[i]; visted {
				break
			}

			line := strings.Trim(data[i], " ")
			if line == "" {
				exited = true
				break
			}

			vistedInstructions[i] = true

			lineSplit := strings.Split(line, " ")
			instruction := lineSplit[0]

			if postion == i && instruction == from {
				instruction = to
			}

			operation := lineSplit[1]
			value, err := strconv.Atoi(operation)
			check(err)

			if instruction == acc {
				accumulator += value
			} else if instruction == jmp {
				i += value - 1
			}
		}
	}

	return exited, accumulator
}

var acc = "acc"
var nop = "nop"
var jmp = "jmp"

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	jmpPositions := findPosition(jmp, data)
	exited, accumulator := runSwap(jmp, nop, data, jmpPositions)

	if exited {
		fmt.Println("Accumilator: ", accumulator)
		return
	}

	nopPositions := findPosition(nop, data)
	exited, accumulator = runSwap(nop, jmp, data, nopPositions)

	fmt.Println("Accumilator: ", accumulator)
	return
}
