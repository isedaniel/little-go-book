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
	fmt.Println(os.Args[1])
}
