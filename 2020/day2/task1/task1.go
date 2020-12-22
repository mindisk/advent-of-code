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

func verifyPassword(rule string, password string) bool {
	min, max, char := parseRule(rule)
	occurences := countCharacters(password, char)

	if occurences >= min && occurences <= max {
		return true
	}
	return false

}

func countCharacters(str string, char byte) int {
	occurences := 0
	for i := 0; i < len(str); i++ {
		if char == str[i] {
			occurences++
		}
	}
	return occurences
}

func parseRule(rule string) (int, int, byte) {
	ruleSplit := strings.Split(rule, " ")
	countSplit := strings.Split(ruleSplit[0], "-")

	min, err := strconv.Atoi(countSplit[0])
	check(err)

	max, err := strconv.Atoi(countSplit[1])
	check(err)

	char := ruleSplit[1][0]

	return min, max, char
}

func main() {
	data, err := ioutil.ReadFile("../input1.txt")
	check(err)

	lines := strings.Split(string(data), "\n")

	allPasswords := 0
	correctPasswords := 0
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], ":")

		// Usually an empty line
		if len(line) < 2 {
			continue
		}

		allPasswords++

		rule := strings.Trim(line[0], " ")
		password := strings.Trim(line[1], " ")
		if verifyPassword(rule, password) {
			correctPasswords++
		}
	}
	fmt.Print("Valid passwords: ", correctPasswords, " out of ", allPasswords)
}
