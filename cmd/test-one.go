package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: cmd \"string to be printed\"")
	}
	fmt.Println(os.Args[1])
}
