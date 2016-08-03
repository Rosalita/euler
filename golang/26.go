package main

import (
  "fmt"
)

func main(){
var numer, denom, result float64 = 1, 2, 0

for i := 0; i < 1000; i++ {
  result = numer / denom
  fmt.Println(result)
  denom += 1
}



}
