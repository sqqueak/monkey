package lexer

import "github.com/sqqueak/monkey/token"

type Lexer struct {
	input		 string
	position	 int
	readPosition int
	ch			 byte
}

// Creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Parse text to get next token
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    // Consume whitespace since Monkey doesn't care about whitespace
    l.skipWhitespace()

    // Consume a character and assign it to a token type
    switch l.ch {
        case '=':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
            } else {
                tok = newToken(token.ASSIGN, l.ch)
            }
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case '-':
            tok = newToken(token.MINUS, l.ch)
        case '!':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
            } else {
                tok = newToken(token.BANG, l.ch)
            }
        case '/':
            tok = newToken(token.SLASH, l.ch)
        case '*':
            tok = newToken(token.ASTERISK, l.ch)
        case '<':
            tok = newToken(token.LT, l.ch)
        case '>':
            tok = newToken(token.GT, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            // Consume tokens which are longer than 1 character
            if isLetter(l.ch) {
                tok.Literal = l.readIdentifier()
                tok.Type = token.LookupIdent(tok.Literal)
                return tok
			} else if isDigit(l.ch) {
                tok.Type = token.INT
                tok.Literal = l.readNumber()
                return tok
            } else {
                tok = newToken(token.ILLEGAL, l.ch)
            }
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

// Reads a single character
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) { // end of line or no character
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1 // consume character
}

// Reads strings of identifiers
func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position] // slice of input containing identifier
}

// Reads blocks of digits
func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position] // slice of input containing number
}

// Skips over blocks of whitespace
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

// Peeks the next character
func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

// Checks if the passed in byte is uppercase or lowercase letter, or an underscore
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Checks if the passed in byte is a digit between 0 and 9.
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
