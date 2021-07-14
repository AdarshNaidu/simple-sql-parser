package sql

import (
	"fmt"
	"io"
)

type SelectStatement struct {
	Fields []string
	TableName string
}

func (s SelectStatement) String() string {
	var fields string

	for _, field := range s.Fields {
		fields += field + " "
	}

	return fmt.Sprintf("Fields: %s\nTable Name: %s\n", fields, s.TableName)
}

type InsertStatement struct {
	Fields []string
	Values []string
	TableName string
}

func (s InsertStatement) String() string {
	var fields string

	for _, field := range s.Fields {
		fields += field + " "
	}

	var values string

	for _, value := range s.Values {
		values += value + " "
	}

	return fmt.Sprintf("Fields: %s\nValues: %s\nTable Name: %s\n", fields, values, s.TableName)
}

type Parser struct {
	s *Scanner
	buf struct {
		tok Token
		lit string
		n int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (interface{}, error) {	

	switch tok, _ := p.scanIgnoreWhitespace(); tok {
	case SELECT:
		p.unscan()
		return p.parseSelect()
	case INSERT:
		p.unscan()
		return p.parseInsert()
	}
	return nil, fmt.Errorf("Sorry, only SELECT and INSERT is supported.")
}

func (p *Parser) parseSelect() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	// first token should be select
	if tok, lit := p.scanIgnoreWhitespace(); tok != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	// after select should be fields or asterisk
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if(tok != IDENT && tok != ASTERISK) {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}

		stmt.Fields = append(stmt.Fields, lit)

		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	// next should be from 
	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	// table name
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	return stmt, nil
}


func (p *Parser) parseInsert() (*InsertStatement, error) {
	stmt := &InsertStatement{}

	// first token should be INSERT
	if tok, lit := p.scanIgnoreWhitespace(); tok != INSERT {
		return nil, fmt.Errorf("found %q, expected INSERT", lit)
	}

	// second token should be into
	if tok, lit := p.scanIgnoreWhitespace(); tok != INTO {
		return nil, fmt.Errorf("found %q, expected INTO", lit)
	}

	// third token table name
	tok, lit := p.scanIgnoreWhitespace();
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	if tok, lit := p.scanIgnoreWhitespace(); tok != OPENING_BRACKET {
		return nil, fmt.Errorf("found %q, expect (", lit)
	}

	for {
		tok, lit := p.scanIgnoreWhitespace()

		if(tok == CLOSING_BRACKET) { 
			p.unscan()
			break 
		}

		if(tok != IDENT) {
			return nil, fmt.Errorf("found %q, expected field name or )", lit)
		}

		stmt.Fields = append(stmt.Fields, lit)

		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CLOSING_BRACKET {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != VALUES {
		return nil, fmt.Errorf("found %q, expected VALUES", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != OPENING_BRACKET {
		return nil, fmt.Errorf("found %q, expect (", lit)
	}
	
	for {
		tok, lit := p.scanIgnoreWhitespace()

		if(tok == CLOSING_BRACKET) { 
			p.unscan()
			break 
		}

		if(tok != IDENT) {
			return nil, fmt.Errorf("found %q, expected value or )", lit)
		}

		stmt.Values = append(stmt.Values, lit)

		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CLOSING_BRACKET {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}

	return stmt, nil
}

// ignore the white space and return the next token
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}