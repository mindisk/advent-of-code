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
	firstPosition, secondPosition, char := parseRule(rule)

	if password[firstPosition-1] == char && password[secondPosition-1] != char {
		return true
	} else if password[firstPosition-1] != char && password[secondPosition-1] == char {
		return true
	}
	return false
}

func parseRule(rule string) (int, int, byte) {
	ruleSplit := strings.Split(rule, " ")
	positionSplit := strings.Split(ruleSplit[0], "-")

	firstPosition, err := strconv.Atoi(positionSplit[0])
	check(err)

	secondPosition, err := strconv.Atoi(positionSplit[1])
	check(err)

	char := ruleSplit[1][0]

	return firstPosition, secondPosition, char
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
