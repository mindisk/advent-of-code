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

func parseLine(line string) (string, map[string]int) {
	lineSplit := strings.Split(line, " bags contain ")

	outterBag := strings.Trim(lineSplit[0], " ")
	var innerBags = make(map[string]int)

	if strings.HasPrefix(lineSplit[1], "no") {
		return outterBag, innerBags
	}

	innerBagsSplit := strings.Split(lineSplit[1], ", ")
	for i := 0; i < len(innerBagsSplit); i++ {
		innerBagSplit := strings.Split(innerBagsSplit[i], " ")
		innerBag := strings.Trim(innerBagSplit[1], " ") + " " + strings.Trim(innerBagSplit[2], " ")
		count, err := strconv.Atoi(strings.Trim(innerBagSplit[0], " "))
		check(err)
		innerBags[innerBag] = count
	}

	return outterBag, innerBags
}

func countBags(initalBags map[string]int, bags map[string]map[string]int) int {
	count := 0
	for initalBag, intialCount := range initalBags {
		foundBags, ok := bags[initalBag]
		if ok {
			count += intialCount * countBags(foundBags, bags)
		}
		count += intialCount
	}
	return count
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	var bags = make(map[string]map[string]int)
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")

		if line == "" {
			continue
		}

		outterBag, innerBags := parseLine(line)

		for innerBag, innerBagCount := range innerBags {
			bgs, outterBagExist := bags[outterBag]
			if outterBagExist {
				innerBgs, innerBagExist := bgs[innerBag]
				if innerBagExist {
					bgs[innerBag] = innerBgs + innerBagCount
				} else {
					bgs[innerBag] = innerBagCount
				}
			} else {
				var dic = make(map[string]int)
				dic[innerBag] = innerBagCount
				bags[outterBag] = dic
			}
		}
	}

	count := 0
	sgBags, ok := bags["shiny gold"]
	if ok {
		count = countBags(sgBags, bags)
	} else {
		fmt.Println("No shiny gold bag found")
	}

	fmt.Println("Bags that shiny gold contain : ", count)
}
