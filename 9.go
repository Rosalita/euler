package main

import (
  "fmt"
  "math"
)

func main(){

for x:= 1; x < 10; x++ {
    for y:=1; y < 9; y++{
      a, b := x, x + y
      if b > 9 {
        continue
      }
      fmt.Printf("a is %d, b is %d\n",a, b)
      a2 := a*a
      b2 := b*b
      fmt.Printf("a^2 is %d, b^2 is %d\n", a2, b2)
      c2 := float64(a2 + b2)
      fmt.Printf("c^2 is %f\n",c2)
      c := (math.Sqrt(c2))
      fmt.Printf("c is %f\n", c)
    }

}

}
