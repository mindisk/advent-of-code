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

func collectData(data []string) []int {
	var set []int
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		check(err)

		set = append(set, value)
	}
	return set
}

func findInvalidNumber(premable int, set []int) int {
	invalidNumber := -1
	for i := premable; i < len(set); i++ {
		number := set[i]
		subset := set[i-premable : i]
		isValid := isNumberValid(number, subset)

		if !isValid {
			invalidNumber = number
			break
		}
	}
	return invalidNumber
}

func isNumberValid(number int, subset []int) bool {
	isValid := false
	for i := 0; i < len(subset); i++ {
		num1 := subset[i]
		for y := 0; y < len(subset); y++ {
			if i == y {
				continue
			}
			num2 := subset[y]

			if number == num1+num2 {
				isValid = true
				break
			}
		}

		if isValid {
			break
		}
	}
	return isValid
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	set := collectData(data)
	invalidNumber := findInvalidNumber(25, set)

	fmt.Println("Invalid number: ", invalidNumber)
}
