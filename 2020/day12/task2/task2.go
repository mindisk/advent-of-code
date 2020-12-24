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
	Waypoint            map[string]int
}

func adjustWaypoint(position Position, instruction Instruction) Position {
	var waypointDirections []string
	var waypointUnits []int
	for direction, units := range position.Waypoint {
		waypointDirections = append(waypointDirections, direction)
		waypointUnits = append(waypointUnits, units)
	}

	clockwise := true
	if instruction.Direction == "R" {
		clockwise = true
	} else if instruction.Direction == "L" {
		clockwise = false
	} else {
		return position
	}

	offset := instruction.Units / 90
	if !clockwise {
		offset = 4 - offset
	}

	if offset == 4 {
		offset = 0
	}

	for i := 0; i < len(waypointDirections); i++ {
		pos := offset + i
		if pos >= 4 {
			pos = pos % 4
		}
		position.Waypoint[waypointDirections[i]] = waypointUnits[pos]
	}
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
	position := Position{N: 0, S: 0, E: 0, W: 0, Degrees: 0, Waypoint: make(map[string]int)}
	position.Waypoint["E"] = 10
	position.Waypoint["N"] = 1
	position.Waypoint["W"] = 0
	position.Waypoint["S"] = 0
	for i := 0; i < len(instructions); i++ {
		position = applyInstruction(position, instructions[i])
	}
	return position
}

func move(position Position, times int) Position {
	for direction, units := range position.Waypoint {
		if direction == "N" {
			position.N += (units * times)
		} else if direction == "E" {
			position.E += (units * times)
		} else if direction == "S" {
			position.S += (units * times)
		} else if direction == "W" {
			position.W += (units * times)
		}
	}
	return position
}

func applyInstruction(position Position, instruction Instruction) Position {
	if instruction.Direction == "N" {
		position.Waypoint["N"] += instruction.Units
	} else if instruction.Direction == "S" {
		position.Waypoint["S"] += instruction.Units
	} else if instruction.Direction == "E" {
		position.Waypoint["E"] += instruction.Units
	} else if instruction.Direction == "W" {
		position.Waypoint["W"] += instruction.Units
	} else if instruction.Direction == "F" {
		position = move(position, instruction.Units)
	} else if instruction.Direction == "R" || instruction.Direction == "L" {
		position = adjustWaypoint(position, instruction)
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
