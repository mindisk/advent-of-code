package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Seat object
type Seat struct {
	Label  string
	Column int
	Row    int
	ID     int
}

// BySeatID sorts by ID
type BySeatID []Seat

func (a BySeatID) Len() int           { return len(a) }
func (a BySeatID) Less(i, j int) bool { return a[i].ID < a[j].ID }
func (a BySeatID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
	var lastPosition byte
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			lastPosition = columns[i]
			if lastPosition == byte('L') {
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
	var lastPosition byte
	for i := 0; i < len(rows); i++ {
		if i == len(rows)-1 {
			lastPosition = rows[i]
			if lastPosition == byte('F') {
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
	var seats []Seat
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

		seats = append(seats, seat)
	}

	sort.Sort(BySeatID(seats))
	fmt.Println("Highest seat in a boarding pass:", highestSeat)

	offset := 78
	for i := 0; i < len(seats); i++ {
		if seats[i].ID != i+offset {
			offset++
			fmt.Println("Missing seat: ", i, "-", seats[i])
		}
	}
}
