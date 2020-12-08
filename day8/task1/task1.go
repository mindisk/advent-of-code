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

var acc = "acc"
var nop = "nop"
var jmp = "jmp"

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	accumulator := 0
	nooperation := 0

	var vistedInstructions = make(map[int]bool)
	for i := 0; i < len(data); i++ {
		if _, visted := vistedInstructions[i]; visted {
			break
		}

		line := strings.Trim(data[i], " ")

		if line == "" {
			break
		}

		lineSplit := strings.Split(line, " ")

		vistedInstructions[i] = true
		instruction := lineSplit[0]
		operation := lineSplit[1]
		value, err := strconv.Atoi(operation)
		check(err)

		if instruction == acc {
			accumulator += value
		} else if instruction == nop {
			nooperation += value
		} else if instruction == jmp {
			i += value - 1
		}
	}

	fmt.Println("Accumilator: ", accumulator)
}
