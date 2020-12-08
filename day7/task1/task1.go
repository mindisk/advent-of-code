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

func findBags(initalBags map[string]bool, bags map[string]map[string]bool) map[string]bool {
	var returnBags = make(map[string]bool)
	for initalBag := range initalBags {
		returnBags[initalBag] = true
		foundBags, ok := bags[initalBag]
		if ok {
			bgs := findBags(foundBags, bags)
			for bg := range bgs {
				returnBags[bg] = true
			}
		}
	}
	return returnBags
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	var uniqueBags map[string]bool
	var bags = make(map[string]map[string]bool)
	for i := 0; i < len(data); i++ {
		line := strings.Trim(data[i], " ")

		if line == "" {
			continue
		}

		outterBag, innerBags := parseLine(line)

		for innerBag := range innerBags {
			value, ok := bags[innerBag]
			if ok {
				value[outterBag] = true
			} else {
				var dic = make(map[string]bool)
				dic[outterBag] = true
				bags[innerBag] = dic
			}
		}
	}

	sgBags, ok := bags["shiny gold"]
	if ok {
		uniqueBags = findBags(sgBags, bags)
	} else {
		fmt.Println("No shiny gold bag found")
	}

	fmt.Println("Bags that contian shiny gold: ", len(uniqueBags))
}
