package main

import (
  "fmt"
)

func main(){
  s:= []int{0}
  for x := 1; x < 100; x++ {
     for y := 1; y < 100; y++ {
       xy := x * y
       fmt.Printf("xy is %v\n", xy)


       for i, v := range s {
         fmt.Printf("v is %v\n", v)
         if xy != s[i] {
           fmt.Printf("%v is not in %v, appending it\n",xy, v)
            s = append(s, xy)
            continue
           }
           continue

       }
     }

  }
}
