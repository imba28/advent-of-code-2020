package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r     *bufio.Reader
	items chan<- Item
}

func (s *Scanner) scan() (Token, string) {
	char := s.read()

	if isWhitespace(char) {
		s.unread()
		return s.scanWhitespace()
	}
	if isLetter(char) {
		s.unread()
		return s.scanLetter()
	}
	if isDigit(char) {
		s.unread()
		return s.scanNumber()
	}

	switch char {
	case eof:
		return EOF, ""
	case ',':
		return COMMA, string(char)
	case '.':
		return DOT, string(char)
	}

	return ILLEGAL, string(char)
}

func (s *Scanner) run() {
	for state := lexIdent; state != nil; {
		state = state(s)
	}
	close(s.items)
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		char := s.read()
		if char == eof {
			break
		}
		if !isWhitespace(char) {
			s.unread()
			break
		}
		buf.WriteRune(char)
	}

	return WS, buf.String()
}

func (s *Scanner) scanNumber() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		char := s.read()
		if char == eof {
			break
		}
		if !isDigit(char) {
			s.unread()
			break
		}
		buf.WriteRune(char)
	}

	return NUMBER, buf.String()
}

func (s *Scanner) scanLetter() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		char := s.read()
		if char == eof {
			break
		}
		if !isLetter(char) && !isDigit(char) {
			s.unread()
			break
		}
		_, _ = buf.WriteRune(char)
	}

	ss := buf.String()
	switch strings.ToLower(ss) {
	case "no":
		return NO, ss
	case "contain":
		return CONTAIN, ss
	}

	return IDENT, ss
}

func (s *Scanner) emit(t Item) {
	s.items <- t
}

func (s *Scanner) scanIgnoreWhitespace() (Token, string) {
	t, literal := s.scan()
	if t == WS {
		t, literal = s.scan()
	}
	return t, literal
}

func NewScanner(r io.Reader, items chan<- Item) *Scanner {
	scanner := &Scanner{
		r:     bufio.NewReader(r),
		items: items,
	}
	go scanner.run()
	return scanner
}
