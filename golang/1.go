package main

import(
  "fmt"
)

func main(){
  x := 0
  total:= 0
  fmt.Println("to find sum of numbers less than x that are multiples of 3 or 5, enter x")
  fmt.Scanln(&x)
  fmt.Printf("x is %d\n", x)
  for i:= 1; i < x; i++ {
    fmt.Printf("i is %d\n", i)
       if i % 3 == 0 || i % 5 == 0 {
         fmt.Printf("%d divides by 3 or 5\n", i)
         total += i
         }
    }
    fmt.Printf("answer is: %d\n", total)
}
