package sql

import (
	"fmt"
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