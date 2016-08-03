package main

import (
	"fmt"
)

func main() {
  var triangle, count, i, j int64
	for  i = 12375; i < 13001; i++ {  // 2079 has 320, 5984 has 480, 765765000 has 576
		triangle = genTriangle(i)
    count = 0
		    for j =1; j <= triangle; j++ {
					 if triangle % j == 0 {
						 count ++
						 if count > 500{
							 fmt.Printf("i= %d Triangle %d has %d divisors\n", i, triangle, count)
							 return
						 }
					 }
					 continue

	      }

	}
	fmt.Println("did not find a suitable triangle")
}

func genTriangle(n int64) int64 {
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
