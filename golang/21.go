package main

import (
	"fmt"
)

func main() {
  amiNums := []int{0} // a slice to store the amicable numbers in
  for i := 1; i < 10001; i ++ {
   y := addNums(getDivs(i))
   x := addNums(getDivs(y))
   if (i == x) && (x != y) {
     fmt.Printf("%d and %d are amicable numbers\n", x, y)
     for _, v := range amiNums { // check the slice of amicable numbers
       if v != x { // if the first amicable number is not already there
          amiNums = append(amiNums, x) // add it to the slice
          break
       }
       if v != y { // if the second amicable number is not already there
         amiNums = append(amiNums, y) // add it to the slice
         break
       }
     }
   }
}
fmt.Println(addNums(amiNums)) // print the sum of the amicable numbers
}

func getDivs(x int) []int{ // function to return the proper divisors for x
	var divs []int // create a new slice to store the divisors in
	for i:= x; i > 0; i -- {
      if (x % i == 0) && (x != i) { // if i divides evenly and is not x
        divs = append(divs, i) // add i to the slice of divs
      }
	}
  return divs // return slice containing all proper divisors
}

func addNums(nums []int) int {
  total := 0 // initialise a total
  for _, v := range nums { // loop through all the proper divisors
    total += v // add each one to the total
  }
  return total
}
