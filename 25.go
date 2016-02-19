package main

import(
  "fmt"
  "strconv"
)

type Fibnum struct {
  i float64  //index
  n float64  // fibonnaci number at index i
  s string   // fibonnaci number at index i in string format
}

func(f *Fibnum) calc(){
  var a, b, i float64 = 0, 1, 1
  for i =1; i <= f.i; i ++ {
    a, b = b, a+b
  }
  f.n = a
  f.s = strconv.FormatFloat(f.n, 'f', -1, 64)
  return
}

func main(){

test := new(Fibnum)
test.i = 1476
test.calc()
fmt.Println(test.n)
fmt.Println(test.s)
fmt.Println(len(test.s))

}
