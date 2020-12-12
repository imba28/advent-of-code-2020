package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type PackageRule struct {
	BagColor string
	NumberOf int
	Contents []*PackageRule
}

func (pr PackageRule) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%q contains %d:\n", pr.BagColor, pr.NumberOf))

	if pr.Contents != nil {
		for i := range pr.Contents {
			buf.WriteString("- " + pr.Contents[i].String())
		}
	}
	return buf.String()
}

type Parser struct {
	items chan Item
}

func (p *Parser) nextItem() Item {
	for {
		select {
		case item := <-p.items:
			return item
		}
	}
}

func NewParser(r io.Reader) *Parser {
	items := make(chan Item, 2)
	NewScanner(r, items)

	p := Parser{
		items: items,
	}

	return &p
}

func (p *Parser) Parse() ([]*PackageRule, error) {
	var rules []*PackageRule
	bagColors := make(map[string]*PackageRule)

	rootRule := &PackageRule{
		NumberOf: 1,
	}
	currentRule := rootRule

	for item := range p.items {
		if item.token == ILLEGAL {
			return nil, errors.New(item.Literal)
		}
		if item.token == EOF {
			break
		}
		if item.token == DOT {
			rootRule.Contents = append(rootRule.Contents, currentRule)
			rules = append(rules, rootRule)
			rootRule = &PackageRule{}
			currentRule = rootRule
		}

		if item.token == CONTAIN || item.token == COMMA {
			if currentRule != rootRule {
				rootRule.Contents = append(rootRule.Contents, currentRule)
			}
			currentRule = &PackageRule{}
		}

		if item.token == IDENT {
			if v, ok := bagColors[item.Literal]; ok {
				if rootRule == currentRule {
					rootRule = v
				}
				currentRule = v
			} else {
				currentRule.BagColor = item.Literal
				bagColors[item.Literal] = currentRule
			}
		}

		if item.token == NUMBER {
			currentRule.NumberOf, _ = strconv.Atoi(item.Literal)
		}
	}

	return rules, nil
}
