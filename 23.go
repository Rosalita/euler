package main

import (
	"fmt"
)

func main() {
	var abundantNums []int
	//  var abundantSums []int
	for i := 1; i < 31; i++ {
		if isAbundant(i) {
			abundantNums = append(abundantNums, i)
		}
	}

	fmt.Println(abundantNums)

	for i1, v1 := range abundantNums {
		fmt.Printf("i1 is %d - v1 is %d\n", i1, v1)

		for i2, v2 := range abundantNums[i1:] { // skip values in second loop that are lower value in first loop
			fmt.Printf("i2 is %d - v2 is %d\n", i2, v2)

			//  nextSum :=  abundantNums[v2] + abundantNums[v1]
			// if nextSum is not in abundantSums already
			//  abundantSums = append(abundantSums, nextSum)
		}
	}
}

func isAbundant(x int) bool {
	divs := getDivs(x)
	sumDivs := addNums(divs)
	if sumDivs > x {
		return true
	}
	return false
}

func getDivs(x int) []int {
	var divs []int
	for i := x; i > 0; i-- {
		if (x%i == 0) && (x != i) {
			divs = append(divs, i)
		}
	}
	return divs
}

func addNums(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
