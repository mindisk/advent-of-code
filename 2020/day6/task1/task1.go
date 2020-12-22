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

		var group = make(map[byte]bool)
		for line != "" {
			for y := 0; y < len(line); y++ {
				group[line[y]] = true
			}

			i++
			line = strings.Trim(data[i], " ")
		}

		count += len(group)
	}

	fmt.Println("Answered yes: ", count)
}
