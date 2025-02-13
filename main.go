package main

import (
	"fmt"
	"os"
	"os/user"
	"supersolik/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to Monkey REPL!\nFeel free to type in expressions\n\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
