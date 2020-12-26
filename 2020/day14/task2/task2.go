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

func parseData(lines []string) map[uint]uint {
	var memory = make(map[uint]uint)

	var mask string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}

		if strings.HasPrefix(line, "mask") {
			maskString := strings.Split(line, " = ")[1]
			mask = buildMaskSet(maskString)
		} else {
			lineSplit := strings.Split(line, " = ")

			val := lineSplit[1]
			u64, err := strconv.ParseUint(val, 10, 32)
			check(err)

			memoryPosString := lineSplit[0][strings.Index(lineSplit[0], "[")+1 : strings.Index(lineSplit[0], "]")]
			memoryPos, err := strconv.ParseUint(memoryPosString, 10, 32)
			check(err)

			positions := applyMask(uint(memoryPos), mask)
			for y := 0; y < len(positions); y++ {
				memory[positions[y]] = uint(u64)
			}
		}
	}
	return memory
}

func applyMask(value uint, mask string) []uint {

	prevActivePos := -1
	var cache [][]uint
	for i := 0; i < len(mask); i++ {
		var posCache []uint
		if mask[i] == '0' {
			cache = append(cache, posCache)
			continue
		}

		if prevActivePos == -1 {
			if mask[i] == '1' {
				posCache = append(posCache, value|(1<<i))
			} else {
				posCache = append(posCache, value|(1<<i))
				posCache = append(posCache, value & ^(1<<i))
			}
			prevActivePos = i
			cache = append(cache, posCache)
		} else {
			prevVals := cache[prevActivePos]

			for y := 0; y < len(prevVals); y++ {
				val := prevVals[y]
				if mask[i] == '1' {
					posCache = append(posCache, val|(1<<i))
				} else {
					posCache = append(posCache, val|(1<<i))
					posCache = append(posCache, val & ^(1<<i))
				}
				prevActivePos = i
			}
			cache = append(cache, posCache)
		}
	}
	return cache[prevActivePos]
}

func buildMaskSet(mask string) string {
	var maskSet []byte
	for i := len(mask) - 1; i >= 0; i-- {
		maskSet = append(maskSet, mask[i])
	}
	return string(maskSet)
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
