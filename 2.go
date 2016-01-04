package main

import (
	"fmt"
)

func main() {
	x, count, total := 0, 0, 0
	fmt.Println("To generate fibonacci for values less than x, please input x:")
	fmt.Scanln(&x)
	for i, j := 0, 1; i < x; i, j = i+j, i {
		if count > 1 {
			if i%2 == 0 {
				total += i
			}
		}
		count += 1
	}
	fmt.Printf("total of even numbers is %d\n", total)
}
