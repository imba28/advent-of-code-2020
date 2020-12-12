package main

var eof = rune(0)

func isWhitespace(char rune) bool {
	return char == ' ' || char == '\n'
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
