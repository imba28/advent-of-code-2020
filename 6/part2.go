package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Group struct {
	GroupAnswers GroupAnswers
	Member       int
}

func (g *Group) ProcessAnswers(answers string) {
	g.Member++
	for _, part := range strings.Split(answers, "") {
		g.GroupAnswers[rune(part[0])]++
	}
}

func (g *Group) CountSameAnswers() int {
	counts := 0
	for _, count := range g.GroupAnswers {
		if count == g.Member {
			counts++
		}
	}
	return counts
}

func (g *Group) Reset() {
	g.Member = 0
	g.GroupAnswers.Reset()
}

func NewGroup() Group {
	return Group{
		GroupAnswers: make(GroupAnswers),
		Member:       0,
	}
}

type GroupAnswers map[rune]int

func (gc *GroupAnswers) Reset() {
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
	g := NewGroup()
	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			sum += g.CountSameAnswers()
			g.Reset()
		} else {
			g.ProcessAnswers(line)
		}
	}

	if len(g.GroupAnswers) > 0 {
		sum += g.CountSameAnswers()
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
