package main

import (
  "fmt"
)

func main(){
  var dict map[int]string
  dict = map[int]string{1:"one", 2:"two", 3:"three", 4:"four", 5:"five"}
  fmt.Printf(" %s %s %s %s %s",dict[1],dict[2],dict[3],dict[4],dict[5])
}
