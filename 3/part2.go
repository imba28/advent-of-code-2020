package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Map [][]rune

func (m Map) CountTrees(stepX, stepY int) int {
	counter, x, y := 0, 0, 0
	for y < len(m) {
		if m[y][x] == '#' {
			counter++
		}
		x = (x + stepX) % len(m[0])
		y += stepY
	}
	return counter
}

func NewMap(filePath string) (Map, error) {
	var field Map

	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		row := make([]rune, len(text))

		for j, token := range text {
			row[j] = token
		}
		field = append(field, row)
	}

	return field, nil
}

func main() {
	field, err := NewMap("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	treeChannel := make(chan int)
	steps := [...][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, step := range steps {
		go func(x, y int) {
			treeChannel <- field.CountTrees(x, y)
		}(step[0], step[1])
	}

	result := -1
	for range steps {
		if result == -1 {
			result = <-treeChannel
		} else {
			result *= <-treeChannel
		}
	}
	fmt.Println(result)
}
