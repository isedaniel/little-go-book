package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.SetFlags(0)
		log.Fatalln("Usage: cmd \"string to be printed\"")
	}
	fmt.Println(Foo())
}

func Foo() string {
	return "I'll print anything I want anyway"
}
