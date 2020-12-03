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

func countTrees(rightStep int, downStep int, data []string) int {
	tree := byte('#')
	empty := byte('.')
	treeCount := 0

	rightPosition := 0
	for i := 0; i < len(data); i += downStep {
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

		rightPosition += rightStep

		// If we reach the edge, move start to the leftmost side.
		if rightPosition > (len(line) - 1) {
			rightPosition = rightPosition - len(line)
		}

		if line[rightPosition] == tree {
			treeCount++
		}
	}
	return treeCount
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	// Right 1, down 1.
	stepRight1 := 1
	stepDown1 := 1
	treeCount1 := countTrees(stepRight1, stepDown1, data)
	fmt.Println("Trees counted for right 1, down 1: ", treeCount1)

	// Right 3, down 1.
	stepRight2 := 3
	stepDown2 := 1
	treeCount2 := countTrees(stepRight2, stepDown2, data)
	fmt.Println("Trees counted for right 3, down 1: ", treeCount2)

	// Right 5, down 1.
	stepRight3 := 5
	stepDown3 := 1
	treeCount3 := countTrees(stepRight3, stepDown3, data)
	fmt.Println("Trees counted for right 5, down 1: ", treeCount3)

	// Right 7, down 1.
	stepRight4 := 7
	stepDown4 := 1
	treeCount4 := countTrees(stepRight4, stepDown4, data)
	fmt.Println("Trees counted for right 7, down 1: ", treeCount4)

	// Right 1, down 1.
	stepRight5 := 1
	stepDown5 := 2
	treeCount5 := countTrees(stepRight5, stepDown5, data)
	fmt.Println("Trees counted for right 1, down 2: ", treeCount5)

	fmt.Println("Traversed trees multiplied:", treeCount1*treeCount2*treeCount3*treeCount4*treeCount5)
}
