package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	error := "Usage: cmd \"string to be printed\""
	message := "I'll print anything i want anyway"
	m := NewMessages(error, message)

	if len(os.Args) != 2 {
		log.SetFlags(0)
		log.Fatalln(m.error)
	}

	fmt.Println(Foo(m.message))
	m.Bar()
	fmt.Println(m.message)

	newerr := newerror{
		m: m,
		error: "this is new error!",
	}
	newerr.Baz()
	fmt.Println(newerr.error)
	fmt.Println(m.message)
}

func Foo(s string) string {
	return s + " (this passed trough Foo function)"
}
