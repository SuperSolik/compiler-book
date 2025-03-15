package repl

import (
	"bufio"
	"fmt"
	"io"
	"supersolik/monkey/eval"
	"supersolik/monkey/lexer"
	"supersolik/monkey/object"
	"supersolik/monkey/parser"
)

const PROMT = "|> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment(nil)

	for {
		fmt.Fprintf(out, PROMT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if evaluated := eval.Eval(program, env); evaluated != nil {
			io.WriteString(out, evaluated.Inspect()+"\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%s\n", msg))
	}
}
