package main

import (
	"fmt"
  "io/ioutil"
  "strings"
)


func main() {
  data, err := ioutil.ReadFile("13.txt")
  if err != nil {
    panic(err)
  }

stringData := string(data)
slicesOfRows := strings.Split(stringData, "\n")

		fmt.Printf("first long number is %s\n", slicesOfRows[0])
    fmt.Printf("second long number is %s\n", slicesOfRows[1])
    fmt.Printf("third long number is %s\n", slicesOfRows[2])

}
