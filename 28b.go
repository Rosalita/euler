package main

import(
  "fmt"
)

func main(){
n:= 5
diag:= make([]bool, (n*n) +1)
  diag[1] = true

counter := 0
mod:= 2
for i, v := range diag{ // run through a slice of bools setting diagonal numbers to true
 if v == true && (i + mod) < (n*n)+1{
   diag[i+mod] = true
    counter++
 }
 if counter == 4{
   mod = mod +2
   counter = 0
 }
  fmt.Println(i,v)
}
 counter = 0
  for i, v := range diag{
   if v == true{
     counter += i
     }
  }
  fmt.Println(counter)

}


// diagonal numbers in 5*5
// 1, 3, 5, 7, 9, 13, 17, 21, 25
//   +2,+2,+2,+2, +4, +4, +4, +4

// diagonal numbers in 7*7
// as above plus  31, 37, 43, 49
//                +6, +6, +6, +6
