package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func numbersFromFile(filePath string) ([]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var numbers []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	numbers, err := numbersFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := range numbers {
		for j := range numbers {
			if j == i {
				continue
			}

			if numbers[i]+numbers[j] == 2020 {
				fmt.Printf("%d * %d = %d", numbers[i], numbers[j], numbers[i]*numbers[j])
				os.Exit(0)
			}
		}
	}

	log.Fatal("something went wrong")
}
