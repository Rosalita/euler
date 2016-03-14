package main

import (
  "fmt"
//  "bytes"
)


func main(){
  money := map[int]string{
    1:"1p",
    2:"2p",
    5:"5p",
    10:"10p",
    20:"20p",
    50:"50p",
    100:"£1",
    200:"£2",
  }

//  var buffer bytes.Buffer
  fmt.Println(money)

  // how many ways can £2 be made from all these coins
// amount := 200
//var combinations []string
//var potential string = ""
for i, v := range money{
fmt.Println(i,v)

//   for n := amount; n > 0; n = n - v {
//     if n == 0 {
//       potential = buffer.String()
//       combinations = append(combinations, potential)
//       continue
//     }
//     buffer.WriteString(money[n])
//     fmt.Println(n)
//     fmt.Println(money[n])
//   }

 }


}
