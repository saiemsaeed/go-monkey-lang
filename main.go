package main

import (
	"os/user"

	"github.com/saiemsaeed/monkey-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	println("Hello " + user.Username + "! Welcome to the Monkey programming language REPL.")

	repl.Start() // Call the REPL function to start the Read-Eval-Print Loop
}
