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

// scan the next token from the scanner
func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case '(':
		return OPENING_BRACKET, string(ch)
	case ')':
		return CLOSING_BRACKET, string(ch)
	}

	return ILLEGAL, string(ch)
}

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

	// check if the string is a reserved word
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	case "INSERT":
		return INSERT, buf.String()
	case "INTO":
		return INTO, buf.String()
	case "VALUES":
		return VALUES, buf.String()
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
