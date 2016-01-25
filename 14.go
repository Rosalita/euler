package main

import (
	"fmt"
)

func main() {
	startChain, longestChain := 0, 0

	for j := 2; j < 1000000; j++ {
		chain := generateChain(j)
		//   fmt.Printf("length of chain is %d\n", len(chain))
		if len(chain) > longestChain {
			longestChain = len(chain)
			startChain = j
		}
	}
	fmt.Printf("longest chain is %d it started on %d\n", longestChain, startChain)
	fmt.Println(generateChain(startChain))
}

func generateNext(n int) int {
	if n%2 == 0 { // if n is even
		return n / 2
	}
	return n*3 + 1
}

func generateChain(n int) []int {
	var chain []int
	chain = append(chain, n)
	if n != 1 {
		for {
			m := generateNext(n)
			chain = append(chain, m)
			if m == 1 {
				break
			}
			n = m
		}

	}
	return chain
}
