package token

type TokenType string

type Token struct {
	Type TokenType  // type of token
	Literal string  // value of token, like "let" or "+" or "else"
}

const ( // all of the symbols that can be in a monkey program
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{ // maps strings to TokenTypes
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	// if something isn't in the list of symbols, then it must be an identifier
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
