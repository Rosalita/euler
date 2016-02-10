package main

import (
	"fmt"
)

func main() {


 //for i:=0;i<10;i++{
	 fmt.Println(getDivs(220))
// }
}

func getDivs(x int) []int{ // returns the proper divisors for x
	var divs []int
	for i:= x; i > 0; i -- {
  fmt.Println(i)
      if x % i == 0{
        divs = append(divs, i)
      }
	}
  return divs
}
