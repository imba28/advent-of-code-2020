package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const preambleLength = 25

func main() {
	numbers := readNumbers("input.txt")

	n := invalidNumber(numbers)
	fmt.Println("invalid number", n)

	set := contiguousSet(numbers, n)
	if set == nil {
		log.Fatal("no series found")
	}
	fmt.Println(set)

	min, max := minMax(set)
	fmt.Printf("%d + %d = %d", min, max, min+max)
}

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

func readNumbers(filePath string) []int {
	c := make(chan int)
	go streamNumbersFromFile(filePath, c)

	var numbers []int
	for n := range c {
		numbers = append(numbers, n)
	}

	return numbers
}

func containsSum(n []int, sum int) bool {
	for i := range n {
		for j := range n {
			if j == i {
				continue
			}

			if n[i]+n[j] == sum {
				return true
			}
		}
	}

	return false
}

func invalidNumber(numbers []int) int {
	slidingWindowLastIdx := preambleLength

	for i := slidingWindowLastIdx; i < len(numbers); i++ {
		if !containsSum(numbers[i-preambleLength:i], numbers[i]) {
			return numbers[i]
		}
	}

	return -1
}

func minMax(n []int) (int, int) {
	var max int = n[0]
	var min int = n[0]

	for _, value := range n {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func contiguousSet(numbers []int, sum int) []int {
	var set []int

	currentSum := 0
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			currentSum += numbers[j]
			if currentSum > sum {
				currentSum = 0
				break
			}
			if j > i && currentSum == sum {
				for z := i; z <= j; z++ {
					set = append(set, numbers[z])
				}
				return set
			}
		}
	}

	return nil
}
