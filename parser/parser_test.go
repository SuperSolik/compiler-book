package parser

import (
	"supersolik/monkey/ast"
	"supersolik/monkey/lexer"
	"testing"
)

func TestParseLetStmts(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram() didn't parse 3 statements, got = %d\n", len(program.Statements))
	}
	expected := []struct {
		name string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range expected {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.name) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedName string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral is not `let`, got=%q\n", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt is not *ast.LetStatement, got=%T\n", stmt)
		return false
	}

	if letStmt.Name.Value != expectedName {
		t.Errorf("letStmt.Name.Value is not %s, got=%s\n", expectedName, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expectedName {
		t.Errorf("letStmt.Name.TokenLiteral() is not %s, got=%s\n", expectedName, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestParseReturnStmts(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram() didn't parse 3 statements, got = %d\n", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement, got=%T\n", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() is not `return`, got=%s\n", returnStmt.TokenLiteral())
		}
	}
}
