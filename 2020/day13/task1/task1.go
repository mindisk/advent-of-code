package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) []int {
	var nums []int
	split := strings.Split(line, ",")
	for i := 0; i < len(split); i++ {
		if split[i] != "x" {
			num, err := strconv.Atoi(split[i])
			check(err)
			nums = append(nums, num)
		}
	}
	return nums
}

func nextBusTime(currentTime int, buses []int) [][]int {
	var nextBusTimeVals [][]int
	for i := 0; i < len(buses); i++ {
		bus := buses[i]

		untilCurrent := currentTime % bus
		nextBusTimeVals = append(nextBusTimeVals, []int{bus, bus - untilCurrent})
	}
	return nextBusTimeVals
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	timestamp, err := strconv.Atoi(data[0])
	check(err)

	buses := parseLine(data[1])

	times := nextBusTime(timestamp, buses)
	earliest := []int{0, math.MaxInt32}
	for i := 0; i < len(times); i++ {
		if earliest[1] > times[i][1] {
			earliest = times[i]
		}
	}

	fmt.Println("Realiest bus and time left: ", earliest)
	fmt.Println("Value: ", earliest[0]*earliest[1])

}
