package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findNumber(turns int, startingNumbers []int) int {
	var spokenNumbers = make(map[int][]int)
	turn := 1
	lastNumber := -1
	for turn <= turns {
		if turn <= len(startingNumbers) {
			number := startingNumbers[turn-1]

			spokenNumbers[number] = []int{turn, -1}
			//fmt.Println("Turn ", turn, ", number: ", number)

			lastNumber = number
		} else {
			if spokenNumbers[lastNumber][1] == -1 {
				number := 0

				if _, ok := spokenNumbers[number]; !ok {
					spokenNumbers[number] = []int{1, -1}
				} else {
					spokenNumbers[number][1] = spokenNumbers[number][0]
					spokenNumbers[number][0] = turn
				}

				//fmt.Println("Turn ", turn, ", number: ", number)
				lastNumber = number
			} else {
				number := spokenNumbers[lastNumber][0] - spokenNumbers[lastNumber][1]
				if _, ok := spokenNumbers[number]; !ok {
					spokenNumbers[number] = []int{turn, -1}
				} else {
					spokenNumbers[number][1] = spokenNumbers[number][0]
					spokenNumbers[number][0] = turn
				}

				//	fmt.Println("Turn ", turn, ", number: ", number)
				lastNumber = number
			}
		}
		turn++
	}
	return lastNumber
}

func main() {
	numbers := []int{19, 0, 5, 1, 10, 13}

	fmt.Println("Task1 - turn: ", 2020, "- number: ", findNumber(2020, numbers))
	fmt.Println("Task2 - turn: ", 30000000, "- number: ", findNumber(30000000, numbers))
}
