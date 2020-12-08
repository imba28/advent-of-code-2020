package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func binaryPartition(pattern string, toLower, toUpper rune, lower, upper int) int {
	for _, char := range pattern {
		mid := (lower + upper) / 2

		if char == toLower {
			upper = mid
		} else if char == toUpper {
			lower = mid
		}
	}

	return upper
}

func row(pattern string) int {
	return binaryPartition(pattern, 'F', 'B', 0, 127)
}

func column(pattern string) int {
	return binaryPartition(pattern, 'L', 'R', 0, 7)
}

func seatID(code string) int {
	row, col := row(code[0:len(code)-3]), column(code[len(code)-3:])
	return row*8 + col
}

func readCodes(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var codes []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	return codes, nil
}

func main() {
	codes, err := readCodes("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	maxSeatID := -1
	for _, code := range codes {
		go func(p string) {
			c <- seatID(p)
		}(code)
	}

	for range codes {
		seatID := <-c
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println(maxSeatID)
}
