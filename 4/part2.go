package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = map[string]Validator{
	"byr": validBirthYear,
	"iyr": validIssueYear,
	"eyr": validExpirationDate,
	"hgt": validHeight,
	"hcl": validHairColor,
	"ecl": validEyeColor,
	"pid": validPassportId,
}

type Validator func(string) bool
type Passport map[string]string

func (p Passport) Valid() bool {
	for key, validator := range requiredFields {
		if _, ok := p[key]; !ok {
			return false
		}
		if false || !validator(p[key]) {
			return false
		}
	}

	return true
}

func valueBetween(v string, lower, upper int) bool {
	i, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return i >= lower && i <= upper
}

func validBirthYear(v string) bool {
	return valueBetween(v, 1920, 2002)
}

func validIssueYear(v string) bool {
	return valueBetween(v, 2010, 2020)
}

func validExpirationDate(v string) bool {
	return valueBetween(v, 2020, 2030)
}

func validHeight(v string) bool {
	if strings.HasSuffix(v, "cm") {
		return valueBetween(v[:len(v)-2], 150, 193)
	}
	if strings.HasSuffix(v, "in") {
		return valueBetween(v[:len(v)-2], 59, 76)
	}
	return false
}

func validHairColor(v string) bool {
	r := regexp.MustCompile("#[\\da-f]{6}")
	return r.Match([]byte(v))
}

func validEyeColor(v string) bool {
	var validEyeColors = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	_, ok := validEyeColors[v]
	return ok
}

func validPassportId(v string) bool {
	r := regexp.MustCompile("^[\\d]{9}$")
	return r.Match([]byte(v))
}

func ValidPassportCount(filePath string) (int, error) {

	f, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	defer f.Close()

	counter := 0
	currentPassport := make(Passport, 8)
	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		if text == "" {
			if currentPassport.Valid() {
				counter++
			}
			currentPassport = make(map[string]string, 8)
			continue
		}

		for _, part := range strings.Split(text, " ") {
			p := strings.Split(part, ":")
			currentPassport[p[0]] = p[1]
		}
	}

	return counter, nil
}

func main() {
	validPassports, err := ValidPassportCount("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(validPassports)
}
