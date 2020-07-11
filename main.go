package main

import (
	"fmt"
	"os"
	"os/user"

	"./repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Monkey v1.0.0 -- written by Omar Garcia\n")
	fmt.Printf("Welcome to the Monkey programming language!\n")
	fmt.Printf("USER: %s\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
