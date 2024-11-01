package main

import (
	"fmt"
	"os"
	"os/user"
	"writing/repl"
)

func main() {
	fmt.Println("BRANCH CHECK")
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the TBD programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
