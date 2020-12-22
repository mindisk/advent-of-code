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
		// for i := 0; i < len(grid); i++ {
		// 	fmt.Println(string(grid[i]))
		// }
		// fmt.Println("==================================================================")
		// fmt.Println("==================================================================")
		// fmt.Println("==================================================================")
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
			if currentSeat == 'L' {
				elements := collectAdjecentValues(i, y, snapshot)

				occupied := 0
				for e := 0; e < len(elements); e++ {
					if elements[e] == '#' {
						occupied++
					}
				}
				if occupied == 0 {
					grid[i][y] = '#'
					changed = true
				}
			} else if currentSeat == '#' {
				elements := collectAdjecentValues(i, y, snapshot)

				occupied := 0
				for e := 0; e < len(elements); e++ {
					if elements[e] == '#' {
						occupied++
					}
				}

				if occupied > 4 {
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

	// top left
	h := row - 1
	v := column - 1
	for h >= 0 && v >= 0 {
		value := grid[h][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h--
		v--
	}

	// top
	h = row - 1
	for h >= 0 {
		value := grid[h][column]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h--
	}

	// top right
	h = row - 1
	v = column + 1
	for h >= 0 && v < len(grid[h]) {
		value := grid[h][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h--
		v++
	}

	// right
	v = column + 1
	for v < len(grid[row]) {
		value := grid[row][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		v++
	}

	// bottom right
	h = row + 1
	v = column + 1
	for h < len(grid) && v < len(grid[h]) {
		value := grid[h][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h++
		v++
	}

	// bottom
	h = row + 1
	for h < len(grid) {
		value := grid[h][column]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h++
	}

	// bottom left
	h = row + 1
	v = column - 1
	for h < len(grid) && v >= 0 {
		value := grid[h][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		h++
		v--
	}

	// left
	v = column - 1
	for v >= 0 {
		value := grid[row][v]
		if value != '.' {
			elements = append(elements, value)
			break
		}
		v--
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
			if grid[i][j] == '#' {
				occupied++
			}
		}
	}
	fmt.Println("Occupied seats: ", occupied)
}
