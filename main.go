package main

import (
	"fmt"
	"os"
	"os/user"

	"entei/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Println("Welcome to Entei, the programming language for everyone!")
	repl.Start(os.Stdin, os.Stdout)
}
