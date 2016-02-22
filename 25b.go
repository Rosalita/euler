package main

import(
  "fmt"
  "math/big"
)

func main(){
a, a2, b  := big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(1))
index := 0
for len(a.String()) < 1000{
  a2.Set(a)
  a.Set(b)
  b.Add(a2, b)
  index ++
}
 fmt.Printf("index of first term to contain 1000 digits is %d", index)
}
