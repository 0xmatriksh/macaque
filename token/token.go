package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// available tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "{"
	RPAREN    = "}"
	LBRACE    = "("
	RBRACE    = ")"

	// Keywords
	FUCNTION = "FUNCTION"
	LET      = "LET"
)
