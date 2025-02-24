package parser

import (
	"fmt"
	"os"
	"strings"
)

var INDENT = 0
var INDENT_INCR = 4

func trace(tag string) string {
	if _, ok := os.LookupEnv("TRACE"); ok {
		fmt.Printf("%s%s\n", strings.Repeat(" ", INDENT), "BEGIN "+tag)
		INDENT += INDENT_INCR
	}
	return tag
}

func untrace(tag string) {
	if _, ok := os.LookupEnv("TRACE"); ok {
		INDENT -= INDENT_INCR
		fmt.Printf("%s%s\n", strings.Repeat(" ", INDENT), "END "+tag)
	}
}
