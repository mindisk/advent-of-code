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

func countCombinations(set []int) int64 {
	var indices = make(map[int]bool)
	workset := []int{0}
	for i := 0; i < len(set); i++ {
		indices[set[i]] = true
		workset = append(workset, set[i])
	}

	var cache = make([]int64, workset[len(workset)-1]+4)
	for i := 0; i < len(workset); i++ {
		currentPos := workset[i]
		currentCount := cache[currentPos]

		if i == len(workset)-1 {
			nextPos := currentPos + 3
			cache[nextPos] = cache[nextPos] + currentCount
			break
		}

		for y := 1; y <= 3; y++ {
			nextPos := currentPos + y
			if _, ok := indices[nextPos]; ok {
				if i == 0 {
					cache[nextPos] = 1
				} else {
					cache[nextPos] = cache[nextPos] + currentCount
				}
			}
		}
	}

	return cache[len(cache)-1]
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	set := collectData(data)
	cominations := countCombinations(set)

	fmt.Println("Combinations: ", cominations)
}
