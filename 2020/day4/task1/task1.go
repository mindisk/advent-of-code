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

func collectPassports(data []string) []map[string]string {
	var passports []map[string]string
	for i := 0; i < len(data); i++ {
		var passport = make(map[string]string)

		line := strings.Trim(data[i], " ")
		for line != "" {
			lineSplit := strings.Split(line, " ")

			for y := 0; y < len(lineSplit); y++ {
				entry := strings.Split(lineSplit[y], ":")
				passport[strings.ToLower(entry[0])] = entry[1]
			}

			i++
			line = strings.Trim(data[i], " ")
		}

		passports = append(passports, passport)
	}
	return passports
}

func isPassportValid(passport map[string]string, fields []string) bool {
	for i := 0; i < len(fields); i++ {
		if _, ok := passport[fields[i]]; !ok {
			return false
		}
	}
	return true
}

var birthYear = "byr"
var issueYear = "iyr"
var expirationYear = "eyr"
var height = "hgt"
var hairColor = "hcl"
var eyeColor = "ecl"
var passportID = "pid"
var countryID = "cid"

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	mandatoryFields := [7]string{birthYear, issueYear, expirationYear, height, hairColor, eyeColor, passportID}
	//optionalFields := [1]string{countryID}

	var fields []string = mandatoryFields[0:7]
	fmt.Println(fields)
	passports := collectPassports(data)

	var validPassportCount int
	for i := 0; i < len(passports); i++ {
		if isPassportValid(passports[i], fields) {
			validPassportCount++
		}
	}

	fmt.Println("Valid passports: ", validPassportCount, " out of ", len(passports))
}
