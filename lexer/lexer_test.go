package lexer

import (
	"supersolik/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){};"

	expectedTokens := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Fatalf("expectedTokens[%d] - token type wrong. expected = %q, got = %q", i, expectedToken.Type, tok.Type)
		}

		if tok.Literal != expectedToken.Literal {
			t.Fatalf("expectedTokens[%d] - literal wrong. expected = %q, got = %q", i, expectedToken.Literal, tok.Literal)
		}
	}
}
