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

	count := 0
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")

		groupCount := 0
		var group = make(map[byte]int)
		for line != "" {
			for y := 0; y < len(line); y++ {
				answerCount, ok := group[line[y]]
				if ok {
					group[line[y]] = answerCount + 1
				} else {
					group[line[y]] = 1
				}
			}
			groupCount++
			i++
			line = strings.Trim(data[i], " ")
		}

		for _, answerCount := range group {
			if answerCount == groupCount {
				count++
			}
		}
	}

	fmt.Println("Answered yes: ", count)
}
