package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	mask        string
	memoryIndex int64
	number      int64
}

type Memory map[int64]int64

func (m Memory) write(index, number int64, mask string) {
	for i := range mask {
		switch mask[i] {
		case '1':
			number |= 1 << i
		case '0':
			number &= ^(1 << i)
		}
	}

	m[index] = number
}

func main() {
	c := make(chan instruction)
	memory := make(Memory)

	go streamData("input.txt", c)

	for instruction := range c {
		memory.write(instruction.memoryIndex, instruction.number, instruction.mask)
	}
	fmt.Println(mapSum(memory))
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func mapSum(m map[int64]int64) int64 {
	var sum int64 = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func streamData(filePath string, c chan<- instruction) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	r := regexp.MustCompile("mem\\[([\\d]+)\\]\\s=\\s([\\d]+)")
	mask := ""
	for scanner.Scan() {
		row := scanner.Text()
		if strings.HasPrefix(row, "mask") {
			mask = reverse(row[7:])
		} else {
			loc := r.FindStringSubmatch(row)
			memoryIndex, err := strconv.ParseInt(loc[1], 10, 64)
			if err != nil {
				panic(err)
			}
			number, err := strconv.ParseInt(loc[2], 10, 64)

			c <- instruction{
				mask, memoryIndex, number,
			}
		}
	}

	close(c)
}
