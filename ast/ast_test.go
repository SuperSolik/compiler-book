package ast

import (
	"supersolik/monkey/token"
	"testing"
)

func TestString(t *testing.T) {

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}

	result := program.String()

	if result != "let myVar = anotherVar;" {
		t.Errorf("program.String() incorrect, got=%q", program.String())
	}
}
