package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Map [][]rune

func (m Map) CountTrees() int {
	counter, x, y := 0, 0, 0
	for y < len(m) {
		if m[y][x] == '#' {
			counter++
		}
		x = (x + 3) % len(m[0])
		y++
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
	fmt.Println(field.CountTrees())
}
