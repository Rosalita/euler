package main

import (
  "fmt"
  "math"
  "sort"
)

func main(){
  var start, end float64 = 2, 5
  var terms []float64

  for i := start; i<= end; i ++{  // loop through all integers
    for j := start; j<= end; j++ { // loop through all powers
      fmt.Println(i, j)
      pow := math.Pow(i,j)
      fmt.Println(pow)
      terms = append(terms, pow)

    }
  }
  sort.Float64s(terms)
  fmt.Println(terms)

  for i, v := range terms{
    deleted := 0
    if i == 0 {
      continue
    }
    fmt.Println(len(terms))
    if deleted > 0 && i == len(terms)-1{
      continue
    }
    if terms[i] == terms[i-1]{
      fmt.Printf("need to delete terms[%d] which is %f\n", i, v )
      terms2 := append(terms[:i], terms[i+1:]...)
      fmt.Println("this is the new slice")
      fmt.Println(terms2)
      fmt.Println(terms)
      deleted ++
    }
    fmt.Println(i, v)
  }
}
