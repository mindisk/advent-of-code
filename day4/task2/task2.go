package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

var birthYear = "byr"
var issueYear = "iyr"
var expirationYear = "eyr"
var height = "hgt"
var hairColor = "hcl"
var eyeColor = "ecl"
var passportID = "pid"
var countryID = "cid"

func isPassportValid(passport map[string]string, fields []string) bool {
	isValid := true
	for i := 0; i < len(fields); i++ {
		value, contains := passport[fields[i]]
		if !contains {
			return false
		}
		isValid = isValid && isFieldValid(fields[i], value)
	}
	return isValid
}

func isFieldValid(field string, value string) bool {
	if field == birthYear || field == issueYear || field == expirationYear {
		if len(value) != 4 {
			return false
		}
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if field == birthYear && year >= 1920 && year <= 2002 {
			return true
		} else if field == issueYear && year >= 2010 && year <= 2020 {
			return true
		} else if field == expirationYear && year >= 2020 && year <= 2030 {
			return true
		}
		return false
	} else if field == height {
		isMetric := strings.HasSuffix(value, "cm")
		isImperial := strings.HasSuffix(value, "in")
		if isMetric || isImperial {
			h, err := strconv.Atoi(value[0 : len(value)-2])
			if err != nil {
				return false
			}

			if isMetric && h >= 150 && h <= 193 {
				return true
			} else if isImperial && h >= 59 && h <= 76 {
				return true
			}
			return false
		}
		return false
	} else if field == hairColor {
		if strings.HasPrefix(value, "#") && len(value) == 7 {
			re := regexp.MustCompile("^[a-fA-F0-9]*$")
			val := value[1:len(value)]
			if re.MatchString(val) {
				return true
			}
			return false
		}
		return false
	} else if field == eyeColor {
		if value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth" {
			return true
		}
		return false
	} else if field == passportID {
		re := regexp.MustCompile("^[0-9]*$")
		if len(value) == 9 && re.MatchString(value) {
			return true
		}
		return false
	}
	return false
}

func main() {
	dat, err := ioutil.ReadFile("../input.txt")
	check(err)

	data := strings.Split(string(dat), "\n")

	mandatoryFields := [7]string{birthYear, issueYear, expirationYear, height, hairColor, eyeColor, passportID}
	//optionalFields := [1]string{countryID}

	var fields []string = mandatoryFields[0:7]
	passports := collectPassports(data)

	var validPassportCount int
	for i := 0; i < len(passports); i++ {
		if isPassportValid(passports[i], fields) {
			validPassportCount++
		}
	}

	fmt.Println("Valid passports: ", validPassportCount, " out of ", len(passports))
}
