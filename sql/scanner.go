package sql

import (
	"bufio"
	"io"
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


