package sql

// every valid token has to be of type Token
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF // end of file
	WS // whitespace

	// Literals
	IDENT // identifiers | fieldname, columnname, tablename, etc.

	// Misc characters
	ASTERISK
	COMMA
	OPENING_BRACKET
	CLOSING_BRACKET

	// Keywords
	SELECT
	FROM
	INSERT
	INTO 
	VALUES
)