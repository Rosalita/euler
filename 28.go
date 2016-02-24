package main

import (
  "fmt"
)

func main(){

genSpiral(5)

}


func genSpiral(n int) [][]int{
grid := [][]int{}

if n%2 == 0 {
  fmt.Println("can only generate spirals for odd numbers")
  return grid
  }
  x, y:= n/2, n/2
  fmt.Println(x, y)
  for i:= 1; i <= n *n; i++{

   fmt.Println(i)
  }

 return grid
}

// spiral coordinates in x,y grid size 0 - 4
// ring 1: 2,2
// ring 2: sideE: 3,2  3,3            sideS: 2,3  1,3            sideW: 1,2  1,1            sideN: 2,1  3,1
// ring 3: sideE: 4,1, 4,2  4,3, 4,4  sideS: 3,4  2,4  1,4  0,4  sideW: 0,3  0,2  0,1  0,0  sideN: 1,0  2,0  3,0  4,0

// for sideE x constant and y changes
// for sideS x changes and y constant
// for sideW x constant and y changes
// for sideN x changes and y constant

// for ring 2 | 3*3 - 1*1 = 8 total numbers in the ring
// for ring 3 | 5*5 - 3*3 = 16 total numbers in the ring

//so potentially a 7*7 spiral will have, 7*7 - 5*5 =  24 total numbers in the outer ring
// and ring 4: sideE: 6,1, 6,2  6,3, 6,4, 6,5  6,6
            // sideS: 5,6  4,6  3,6  2,6  1,6  0,6
            // sideW: 0,5  0,4  0,3  0,2  0,1  0,0
            // sideN: 1,0  2,0  3,0  4,0  5,0  6,0

// formula: total of numbers in outer ring of n*n grid |  r =  n² - ((n-2))²
             // n² - (n-2)(n-2)
             // n² - n(n-2) -2(n-2)
             // n² - n² -2n -2n +4
             // -r = -4n +4
             // r = 4n -4
