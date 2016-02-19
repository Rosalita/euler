package main

import(
"fmt"
  "math/big"
)



func main(){
  //n := 5
a, a2, b  := big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(1))

a2.Set(a)
a.Set(b)
b.Add(a2, b)
fmt.Println(a)



//var a, b, i float64 = 0, 1, 1
//for i =1; i <= f.i; i ++ {
//  a, b = b, a + b
//a2 = a
//a = b
//b = a2 + b



//}
//f.n = a
//f.s = strconv.FormatFloat(f.n, 'f', -1, 64)
//return

}

//func fib(){

//}
