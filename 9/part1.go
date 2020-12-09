package main

import (
	"bufio"
	"os"
	"strconv"
)

const preambleLength = 25

func streamNumbersFromFile(filePath string, c chan<- int) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	defer close(c)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		c <- number
	}
}

func containsSum(numbers []int, sum int) bool {
	for i := range numbers {
		for j := range numbers {
			if j == i {
				continue
			}

			if numbers[i]+numbers[j] == sum {
				return true
			}
		}
	}

	return false
}

func main() {
	c := make(chan int)
	go streamNumbersFromFile("input.txt", c)

	var numbers []int

	for number := range c {
		if len(numbers) == preambleLength {
			if !containsSum(numbers, number) {
				println("invalid series", number)
				return
			}
		}

		numbers = append(numbers, number)
		if len(numbers) > preambleLength {
			numbers = numbers[1:]
		}
	}

	println("numbers series is valid")
}
