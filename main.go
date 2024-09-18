package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming Language\n", user.Username)
	fmt.Printf("Enter in any command\n")
	repl.Start(os.Stdin, os.Stdout)
}