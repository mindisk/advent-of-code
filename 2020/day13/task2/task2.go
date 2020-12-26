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

func parseLine(line string) ([]int, []int) {
	var value []int
	var position []int
	split := strings.Split(line, ",")
	for i := 0; i < len(split); i++ {
		if split[i] != "x" {
			num, err := strconv.Atoi(split[i])
			check(err)
			value = append(value, num)
			position = append(position, i)
		}
	}
	return value, position
}

// Calculate remainders and mods based on the value position.
func calculateRemaindersMods(value, position []int) ([]int, []int) {
	max := position[len(position)-1]
	var remainders []int
	for i := 0; i < len(position); i++ {
		remainders = append(remainders, max-position[i])
	}
	return remainders, value
}

// Chinese Remainder Theorem
func chineseRemainder(remainders []int, mods []int) int {
	N := 1
	for i := 0; i < len(mods); i++ {
		N *= mods[i]
	}

	sum := 0
	for i := 0; i < len(mods); i++ {
		bi := remainders[i]
		ni := mods[i]
		Ni := N / ni
		xi := modInverse(Ni, ni)
		sum += bi * Ni * xi
	}

	return sum % N
}

func modInverse(v, m int) int {
	v = v % m
	for x := 1; x <= m; x++ {
		if (v*x)%m == 1 {
			return x
		}
	}
	return 1
}

// Used Chinese Remainder calculator to test example data values
// https://www.dcode.fr/chinese-remainder
func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	value, position := parseLine(data[1])
	r, m := calculateRemaindersMods(value, position)
	time := chineseRemainder(r, m)
	fmt.Println("Earliest bus time: ", time-position[len(position)-1])
}
