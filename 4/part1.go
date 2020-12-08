package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var requiredFields = [...]string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}

type Passport map[string]string

func (p Passport) Valid() bool {
	for _, key := range requiredFields {
		if _, ok := p[key]; !ok {
			return false
		}
	}

	return true
}

func ValidPassportCount(filePath string) (int, error) {
	counter := 0
	f, err := os.Open("input.txt")
	if err != nil {
		return counter, err
	}
	defer f.Close()

	currentPassport := make(Passport, 8)

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		if text == "" {
			if currentPassport.Valid() {
				counter++
			}
			currentPassport = make(map[string]string, 8)
		} else {
			for _, part := range strings.Split(text, " ") {
				p := strings.Split(part, ":")
				currentPassport[p[0]] = p[1]
			}
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
