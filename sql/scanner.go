package sql

import (
	"bufio"
	"io"
)

// lexical scanner object
type Scanner struct {
	R *bufio.Reader
}

// scanner constructor, can take input from any io.Reader interface
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

