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

func parseData(lines []string) map[int]uint {
	var memory = make(map[int]uint)

	var maskMap = make(map[int]bool)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}

		if strings.HasPrefix(line, "mask") {
			mask := strings.Split(line, " = ")[1]
			maskMap = buildMaskMap(mask)
		} else {
			lineSplit := strings.Split(line, " = ")

			val := lineSplit[1]
			u64, err := strconv.ParseUint(val, 10, 32)
			check(err)

			memoryPosString := lineSplit[0][strings.Index(lineSplit[0], "[")+1 : strings.Index(lineSplit[0], "]")]
			memoryPos, err := strconv.Atoi(memoryPosString)
			check(err)

			value := applyMask(uint(u64), maskMap)
			memory[memoryPos] = value
		}
	}
	return memory
}

func applyMask(value uint, maskMap map[int]bool) uint {
	for pos, val := range maskMap {
		if val {
			value |= (1 << pos)
		} else {
			value = value & ^(1 << pos)
		}
	}
	return value
}

func buildMaskMap(mask string) map[int]bool {
	var maskMap = make(map[int]bool)
	for i := len(mask) - 1; i >= 0; i-- {
		val := mask[i]
		if val != 'X' {
			pos := (len(mask) - 1) - i
			value := false
			if val == '1' {
				value = true
			}
			maskMap[pos] = value
		}
	}
	return maskMap
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	memory := parseData(data)
	var sum uint
	for _, value := range memory {
		sum += value
	}

	fmt.Println("Sum: ", sum)

}
