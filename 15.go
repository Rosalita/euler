package main

import (
	"fmt"
)

func main() {
	x, y := 0, 0
	fmt.Println("what is the value of x?")
	fmt.Scanln(&x)
	fmt.Println("what is the value of y?")
	fmt.Scanln(&y)

	fmt.Printf("x is %d and y is %d\n", x, y)

	journeyLen := x + y

	//  for i:= 0; i < journeyLen; i++ {

	//  if journeyLen = 1 // could have moved
	// 1x
	// 1y
	// 2 paths

	//     if journeyLen  = 2 // could have moved
	// 1x + 1y
	// 1y + 1x
	// 2x
	// 2y
	// 4 paths

	//   if journeyLen = 3 // could have moved
	// 2x + 1y
	// 1x + 1y +1x
	// 1x + 2y
	// 1y + 2x
	// 1y + 1x + 1y
	// 2y + 1x
	// 6 paths

	//   if journeyLen = 4 // could have moved
	// 2x + 2y
	// 2y + 2x
	// 1x + 1y + 1x + 1y
	// 1y + 1x + 1y + 1x
	// 1x + 2y + 1x
	// 1y + 2x + 1y
	// 6 paths

	for i := x; i >= 0; i-- {
		fmt.Printf("x is %d\n", i)
		journeyLeft := journeyLen - x

		fmt.Printf("journey left is %d\n", journeyLeft)

	}

}
