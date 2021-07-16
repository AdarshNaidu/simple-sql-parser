package sql

import (
	"fmt"
)

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

	tok, lit = p.scanIgnoreWhitespace()

	if tok == OPENING_BRACKET {
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

	} else if tok == VALUES {
		p.unscan()
	} else {
		return nil, fmt.Errorf("found %q, expected ( or VALUES", lit)
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