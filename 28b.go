package main

import(
  "fmt"
)

func main(){
n:= 1001
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
}
 counter = 0 // find the sum of all the true values
  for i, v := range diag{
   if v == true{
     counter += i
     }
  }
  fmt.Println(counter)
}
