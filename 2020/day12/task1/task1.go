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

// Position of the element
type Position struct {
	N, S, E, W, Degrees int
	Direction           string
}

func adjustPositionDirection(position Position, instruction Instruction) Position {
	clockwise := true
	if instruction.Direction == "R" {
		clockwise = true
	} else if instruction.Direction == "L" {
		clockwise = false
	} else {
		return position
	}

	degrees := position.Degrees
	if clockwise {
		degrees += instruction.Units
	} else {
		degrees -= instruction.Units
	}

	if degrees == 360 {
		degrees = 0
	} else if degrees > 360 {
		degrees = degrees % 360
	} else if degrees < 0 {
		degrees = 360 + degrees
	}

	// adjust direction based on degree
	var direction string
	if degrees == 0 {
		direction = "E"
	} else if degrees == 90 {
		direction = "S"
	} else if degrees == 180 {
		direction = "W"
	} else {
		direction = "N"
	}

	position.Degrees = degrees
	position.Direction = direction
	return position
}

// Instruction to execute
type Instruction struct {
	Direction string
	Units     int
}

func parseData(data []string) []Instruction {
	var instructions []Instruction
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")
		if line == "" {
			continue
		}
		dir := line[0]

		units, err := strconv.Atoi(line[1:])
		check(err)

		instructions = append(instructions, Instruction{Direction: string(dir), Units: units})
	}
	return instructions
}

func applyInstructions(instructions []Instruction) Position {
	position := Position{N: 0, S: 0, E: 0, W: 0, Degrees: 0, Direction: "E"}
	for i := 0; i < len(instructions); i++ {
		position = applyInstruction(position, instructions[i])
	}
	return position
}

func applyInstruction(position Position, instruction Instruction) Position {
	if instruction.Direction == "N" {
		position.N += instruction.Units
	} else if instruction.Direction == "S" {
		position.S += instruction.Units
	} else if instruction.Direction == "E" {
		position.E += instruction.Units
	} else if instruction.Direction == "W" {
		position.W += instruction.Units
	} else if instruction.Direction == "F" {
		position = applyInstruction(position, Instruction{Direction: position.Direction, Units: instruction.Units})
	} else if instruction.Direction == "R" || instruction.Direction == "L" {
		position = adjustPositionDirection(position, instruction)
	}

	return position
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")
	instructions := parseData(data)
	position := applyInstructions(instructions)

	fmt.Println("Position: ", position)
	dist := abs(position.E-position.W) + abs(position.N-position.S)
	fmt.Println("Distance: ", dist)

}
