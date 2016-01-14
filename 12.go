package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d: triangle %d\n", i, genTriangle(i))
	}
}

func genTriangle(n int) int {
	total := n
	for i := n; i > 0; i-- {
		if total > i {
			total = total + i
			continue
		}
		total = i
	}
	return total
}
