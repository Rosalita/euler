package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y int64 = 20, 20
	n := x + y
	paths := new(big.Int)                                  // create a new big.Int value using package math/big
	fmt.Printf("There are %d paths", paths.Binomial(n, x)) // Binomial() method sets value to the binomial coefficient of (n,k)
}
