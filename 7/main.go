package main

import (
	"fmt"
	"log"
	"os"
)

func containsBagColor(rule *PackageRule, bagColor string) bool {
	if rule.BagColor == bagColor {
		return true
	}

	if rule.Contents != nil {
		for i := range rule.Contents {
			if containsBagColor(rule.Contents[i], bagColor) {
				return true
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parser := NewParser(file)
	rules, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	counter := -1
	for i := range rules {
		if containsBagColor(rules[i], "shinygold") {
			counter++
		}
	}

	fmt.Println(counter)
}
