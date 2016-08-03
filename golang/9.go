package main

import (
	"fmt"
	"math"
)

func main() {

	for x := 1; x < 1000; x++ {
		for y := 1; y < 999; y++ {
			a, b := x, x+y
			if b > 999 {
				continue
			}
			//      fmt.Printf("a is %d, b is %d\n",a, b)
			a2 := a * a
			b2 := b * b
			//    fmt.Printf("a^2 is %d, b^2 is %d\n", a2, b2)
			c2 := float64(a2 + b2)
			//    fmt.Printf("c^2 is %f\n",c2)
			c := (math.Sqrt(c2))
			//  fmt.Printf("c is %f\n", c)
			if hasNoDecimals(c) == true {
				if a+b+int(c) == 1000 {
					fmt.Printf("a = %d, b = %d, c = %d | a^2 = %d, b^2 = %d, c^2 = %d\n", a, b, int(c), a2, b2, int(c2))
					fmt.Printf("a + b + c = %d\n", a+b+int(c))
					fmt.Printf("a * b * c = %d\n", a*b*int(c))
				}
			}
		}

	}
}

func hasNoDecimals(value float64) bool {
	value = value * 10
	extra := value - float64(int(value))
	if extra == 0 {
		return true
	} else {
		return false
	}
}
