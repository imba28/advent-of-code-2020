package main

import (
	"bytes"
	"fmt"
)

type stateFn func(*Scanner) stateFn

func lexMisc(s *Scanner) stateFn {
	t, literal := s.scanIgnoreWhitespace()
	item := Item{
		token:   t,
		Literal: literal,
	}

	if t == EOF {
		s.emit(item)
		return nil
	}

	if t == DOT {
		s.emit(item)
		return lexIdent
	}

	if t == CONTAIN || t == COMMA {
		s.emit(item)
		return lexNumber
	}

	return s.errorf("found %q, expected ',' 'contains' or eof", literal)
}

func lexToDot(p *Scanner) stateFn {
	for {
		t, literal := p.scanIgnoreWhitespace()

		if t == EOF {
			return p.errorf("unexpected eof, expected '.'")
		}

		if t == DOT {
			p.emit(Item{
				token:   DOT,
				Literal: literal,
			})
			return lexIdent
		}
	}
}

func lexNumber(p *Scanner) stateFn {
	t, literal := p.scanIgnoreWhitespace()
	if t == NUMBER {
		p.emit(Item{
			token:   NUMBER,
			Literal: literal,
		})
		return lexIdent
	}

	if t == NO {
		p.emit(Item{
			t,
			literal,
		})
		return lexToDot
	}

	return p.errorf("found %q, expected number", literal)
}

func lexIdent(s *Scanner) stateFn {
	var buf bytes.Buffer

	for i := 0; i < 3; i++ {
		t, literal := s.scanIgnoreWhitespace()
		if t == EOF {
			return nil
		}
		if t != IDENT {
			return s.errorf("expected an identifier (%d), but got %q", i, literal)
		}

		if i < 2 {
			buf.WriteString(literal)
		}
	}

	s.emit(Item{
		token:   IDENT,
		Literal: buf.String(),
	})

	return lexMisc
}

func (s *Scanner) errorf(format string, args ...interface{}) stateFn {
	s.emit(Item{
		token:   ILLEGAL,
		Literal: fmt.Sprintf(format, args...),
	})
	return nil
}
