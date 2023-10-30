package main

import "fmt"

func arrayCap() {
	scores := make([]int, 0, 10)
	scores = scores[0:8]
	scores[7] = 42
	fmt.Println(scores)
}
