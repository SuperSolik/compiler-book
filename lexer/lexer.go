package lexer

import (
	"supersolik/monkey/token"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.ReadChar()
	return lexer
}

func (l *Lexer) ReadChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func NewMultiCharToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.ReadChar()
			tok = NewMultiCharToken(token.EQ, "==")
		} else {
			tok = NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = NewToken(token.PLUS, l.ch)
	case '-':
		tok = NewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.ReadChar()
			tok = NewMultiCharToken(token.NOT_EQ, "!=")
		} else {
			tok = NewToken(token.BANG, l.ch)
		}
	case '*':
		tok = NewToken(token.ASTERISK, l.ch)
	case '/':
		tok = NewToken(token.SLASH, l.ch)
	case '<':
		tok = NewToken(token.LT, l.ch)
	case '>':
		tok = NewToken(token.GT, l.ch)
	case ';':
		tok = NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = NewToken(token.LPAREN, l.ch)
	case ')':
		tok = NewToken(token.RPAREN, l.ch)
	case '{':
		tok = NewToken(token.LBRACE, l.ch)
	case '}':
		tok = NewToken(token.RBRACE, l.ch)
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = NewToken(token.ILLEGAL, l.ch)
	}

	l.ReadChar()
	return tok
}

func (l *Lexer) readIdent() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.ReadChar()
	}

	return l.input[pos:l.pos]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.ReadChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	pos := l.pos + 1
	for {
		l.ReadChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}
