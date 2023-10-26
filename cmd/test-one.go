package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	m := messages{
		error: "Usage: cmd \"string to be printed\"",
		message: "I'll print anything i want anyway",
	}
	if len(os.Args) != 2 {
		log.SetFlags(0)
		log.Fatalln(m.error)
	}
	fmt.Println(Foo(m.message))
}

func Foo(s string) string {
	return s + " (this passed trough Foo function)"
}
