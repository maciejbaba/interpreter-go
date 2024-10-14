package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL     = "ILLEGAL"
	END_OF_FILE = "END_OF_FILE"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y - for example name of the function
	INTEGER    = "INTEGER"    // 12334 - just a integer number

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
  COLON = ":"

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
