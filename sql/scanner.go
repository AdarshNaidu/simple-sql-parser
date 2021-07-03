package sql

import (
	"bufio"
	"io"
	"bytes"
	"strings"
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

func (s *Scanner) scanIdent() (tok Token, lit string) {

	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	// check if the string is a reserverd word
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}

	return IDENT, buf.String()
}

func isWhitespace (ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter (ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit (ch rune) bool {
	return ch >= '0' && ch <= '9'
}
