package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Seat object
type Seat struct {
	Label  string
	Column int
	Row    int
	ID     int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSeat(seatLabel string) Seat {
	rows := seatLabel[0:7]
	row := parseRows(rows)

	columns := seatLabel[7:10]
	column := parseColumns(columns)

	return Seat{Label: seatLabel, Row: row, Column: column, ID: (row * 8) + column}
}

func parseColumns(columns string) int {
	column := 0
	lo := 0
	hi := 7
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			if columns[i] == byte('L') {
				column = lo
			} else {
				column = hi
			}
			break
		}

		if columns[i] == byte('L') {
			hi = (hi - (hi-lo)/2) - 1
		} else {
			lo = (lo + (hi-lo)/2) + 1
		}
	}
	return column
}

func parseRows(rows string) int {
	row := 0
	lo := 0
	hi := 127
	for i := 0; i < len(rows); i++ {
		if i == len(rows)-1 {
			if rows[i] == byte('F') {
				row = lo
			} else {
				row = hi
			}
			break
		}

		if rows[i] == byte('F') {
			hi = (hi - (hi-lo)/2) - 1
		} else {
			lo = (lo + (hi-lo)/2) + 1
		}
	}
	return row
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")
	var highestSeat Seat
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")

		if line == "" {
			continue
		}

		seat := getSeat(line)
		if highestSeat.ID <= seat.ID {
			highestSeat = seat
		}
	}

	if highestSeat.Label == "" {
		fmt.Println("Given data has no seats")
	}

	fmt.Println("Highest seat in a boarding pass:", highestSeat.ID)
}
