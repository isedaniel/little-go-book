package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// "fmt"
// "log"
// "os"

func main() {
	scores := make([]int, 10)
	for i := 0; i < 10; i++ {
		scores[i] = int(rand.Intn(100))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst[2:], scores[:1])
	fmt.Println(worst)
	fmt.Println(scores)
}
