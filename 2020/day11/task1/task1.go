package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func collectData(data []string) [][]byte {
	var seats [][]byte
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")
		if line == "" {
			continue
		}
		row := []byte(line)
		seats = append(seats, row)
	}
	return seats
}

func takeSnapshot(grid [][]byte) [][]byte {
	var snapshot [][]byte
	for i := 0; i < len(grid); i++ {
		var rowSnapshot = make([]byte, len(grid[i]))
		copy(rowSnapshot, grid[i])
		snapshot = append(snapshot, rowSnapshot)
	}
	return snapshot
}

func processGrid(grid [][]byte) (bool, [][]byte) {
	changed, grid := fillGrid(grid)
	for changed {
		changed, grid = fillGrid(grid)
	}
	return true, grid
}

func fillGrid(grid [][]byte) (bool, [][]byte) {
	snapshot := takeSnapshot(grid)
	changed := false

	for i := 0; i < len(snapshot); i++ {
		for y := 0; y < len(snapshot[i]); y++ {
			currentSeat := snapshot[i][y]
			if currentSeat == byte('L') {
				elements := collectAdjecentValues(i, y, snapshot)

				occupied := 0
				for e := 0; e < len(elements); e++ {
					if elements[e] == byte('#') {
						occupied++
					}
				}
				if occupied == 0 {
					grid[i][y] = '#'
					changed = true
				}
			} else if currentSeat == byte('#') {
				elements := collectAdjecentValues(i, y, snapshot)

				occupied := 0
				for e := 0; e < len(elements); e++ {
					if elements[e] == byte('#') {
						occupied++
					}
				}

				if occupied >= 4 {
					grid[i][y] = 'L'
					changed = true
				}
			}
		}
	}
	return changed, grid
}

func collectAdjecentValues(row int, column int, grid [][]byte) []byte {
	var elements []byte
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[i]) && !(i == row && j == column) {
				element := grid[i][j]
				elements = append(elements, element)
			}
		}
	}
	return elements
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")
	seats := collectData(data)

	_, grid := processGrid(seats)

	occupied := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == byte('#') {
				occupied++
			}
		}
	}
	fmt.Println("Occupied seats: ", occupied)
}
