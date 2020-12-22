package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
		sort.Ints(set)
	}
	return set
}

func parseAdapters(set []int) (int, int, int) {
	oneJolt := 0
	twoJolt := 0
	threeJolt := 0

	currentAdapterJolt := 0
	for i := 0; i < len(set); i++ {
		joltDiff := set[i] - currentAdapterJolt
		if joltDiff == 1 {
			oneJolt++
		} else if joltDiff == 2 {
			twoJolt++
		} else if joltDiff == 3 {
			threeJolt++
		} else {
			fmt.Println("Jolt error. Diff ", joltDiff, " [", currentAdapterJolt, " ", set[i], " Position: ", i)
			break
		}
		currentAdapterJolt = set[i]
	}
	threeJolt++
	return oneJolt, twoJolt, threeJolt
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	set := collectData(data)
	oneJolt, _, threeJolt := parseAdapters(set)

	fmt.Println("Jolt differences: ", oneJolt*threeJolt)
}
