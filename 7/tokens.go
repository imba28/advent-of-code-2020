package main

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT
	NO
	NUMBER

	// misc characters
	COMMA
	CONTAIN
	DOT
)

type Item struct {
	token   Token
	Literal string
}
