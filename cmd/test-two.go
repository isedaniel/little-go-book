package main

import "fmt"

func arrayCap() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		fmt.Println("i: ", i, "cap: ", cap(scores))
	}
}
