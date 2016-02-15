package main

import (
	"fmt"
)

func main() {
	var abundantNums []int
  var abundantSums []int
	for i := 1; i < 28124; i++ { // check each of the numbers 1 - 28123
		if isAbundant(i) {
			abundantNums = append(abundantNums, i)
		}
	}

//now we have all the abundantNums, populate a slice with all possible abundantSums
	for i1, _ := range abundantNums {
		for i2, _ := range abundantNums[i1:] { // skip values in second loop that are lower value in first loop
			 nextSum := abundantNums[i1] +  abundantNums[i2]
				 if len(abundantSums) == 0{
				 abundantSums = append(abundantSums, nextSum)
			 } else {
          for i, _ := range abundantSums{
						if nextSum == abundantSums[i] { // if nextSum is in the slice of sums
							break // break and do nothing
							}
						if i == len(abundantSums) -1	{ // if at end of slice and nextSum not found in the slice
							abundantSums = append(abundantSums, nextSum) // add nextSum to the slice
						}
					}
			 }
		}
	}

total := 0 // initialise a total to count
// find all integers less than 28123 which are not in the abundantSums slice
for i:= 1; i < 28124; i++ { // for each integer less than 28123
	for j, v := range abundantSums{ // check if its is in abundantSums
		if v == i { // and if it is, break
			break
		}
		if j == len(abundantSums)-1 { //if reached the end of the slice and not found i
						total += i //add it to the total
		}
	}
}
fmt.Printf(" %d is the total sum of positive integers which cannot be written as the sum of two abundant numbers\n", total)

// add all these ints together

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
