package sql

import (
	"bufio"
	"io"
	"bytes"
)

const eof = rune(0)

// lexical scanner object
type Scanner struct {
	r *bufio.Reader
}

// scanner constructor, can take input from any io.Reader interface
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

// read and return the next rune
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) scanWhitespace() (Token, string) {

	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func isWhitespace (ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

