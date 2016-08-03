package main

import (
	"fmt"
)

func main() {

	sos := sumOfSquares(100)
	fmt.Printf("sum of squares is %d\n", sos)
	ss := sumSquared(100)
	fmt.Printf("sum of numbers, squared is %d\n", ss)
	difference := ss - sos
	fmt.Printf("difference between sum of squares and sum squared is %d\n", difference)
}

func sumOfSquares(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		square := (i + 1) * (i + 1)
		total = total + square
	}
	return total
}

func sumSquared(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total = total + (i + 1)
	}
	square := total * total
	return square
}
