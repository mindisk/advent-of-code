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

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	tree := byte('#')
	empty := byte('.')
	treeCount := 0

	rightPosition := 0
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")

		if len(line) == 0 {
			continue
		}

		// Find first position to start with.
		if i == 0 {
			for line[rightPosition] != empty {
				rightPosition++
			}
			continue
		}

		rightPosition += 3

		// If we reach the edge, move start to the leftmost side.
		if rightPosition > (len(line) - 1) {
			rightPosition = rightPosition - len(line)
		}

		if line[rightPosition] == tree {
			treeCount++
		}
	}

	fmt.Println("Trees counted: ", treeCount)
}
