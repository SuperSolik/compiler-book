package repl

import (
	"bufio"
	"fmt"
	"io"
	"supersolik/monkey/lexer"
	"supersolik/monkey/token"
)

const PROMT = "|> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMT)
		if scanner.Scan() {
			line := scanner.Text()
			l := lexer.New(line)

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Fprintf(out, "%+v\n", tok)
			}
		}
	}
}
