package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type GroupCounter map[rune]int

func (gc *GroupCounter) Clear() {
	for key := range *gc {
		delete(*gc, key)
	}
}

func SumOfCounts(filePath string) (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sum := 0
	group := make(GroupCounter)
	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			sum += len(group)
			group.Clear()
		} else {
			for _, part := range strings.Split(line, "") {
				group[rune(part[0])]++
			}
		}
	}

	if len(group) > 0 {
		sum += len(group)
	}

	return sum, nil
}

func main() {
	sum, err := SumOfCounts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
