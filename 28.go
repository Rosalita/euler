package main

import (
  "fmt"
)

func main(){

genSpiral(5)

}


func genSpiral(n int) [][]int{

grid := make([][]int, n) // make a slice called grid to hold a slice for each row
  for i := range grid { // for each item in grid slice
    grid[i] = make([]int, n) // make a slice to represent the row
  }

if n%2 == 0 {
  fmt.Println("can only generate spirals for odd numbers")
  return grid // return an empty grid
  }

totalrings := n/2 +1
fmt.Printf("There are %d rings in total\n", totalrings)

// just set the middle value which is the first ring as its always 1
x, y := n/2, n/2  // x,y is coords of the central point in the spiral
grid[x][y] = 1
next := 2
 printGrid(grid)
println()
m := len(grid)/2
for ring:=2;ring <= totalrings; ring++ {
      for i:= 0; i < ring; i ++{  // sideE: 2,3  3,3
      grid[m+(ring-1)][m+i] = next
      next++
    }
  //  for i:= 0; i < ring; i ++{  // sideS: 3,2  3,1
  //    grid[m+(ring-1)][m-i] = next
  //    next++
  //  }
  //  for i:= 0; i < ring; i ++{  // sideW: 2,1  1,1
  //    grid[m+(ring-1)][m-i] = next
  //    next++
  //  }
  //  for i:= 0; i < ring; i ++{  // sideN: 1,2  1,3
  //    grid[m+(ring-1)][m-i] = next
  //    next++
  //  }
  }
 printGrid(grid)

return grid
}

func printGrid(grid [][]int){
  for row := 0; row < len(grid); row ++{
    for col:=0; col < len(grid); col ++{
      fmt.Printf(" %d",grid[row][col])
    }
    println()
  }

}

// spiral coordinates in x,y grid size 0 - 4
// ring 1: 2,2
// ring 2: sideE: 2,3  3,3            sideS: 3,2  3,1            sideW: 2,1  1,1            sideN: 1,2  1,3
// ring 3: sideE: 1,4, 2,4, 3,4, 4,4  sideS: 4,3  4,2  4,1  4,0  sideW: 3,0  2,0  1,0  0,0  sideN: 0,1  0,2  0,3  0,4

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

// formula: total number of rings (r) from n
//  n = 9 | r = 5 | d -4   9 / 2 = 4.5 = int divide 9/2 = 4, then + 1
//  n = 7 | r = 4 | d -3   7 / 2 = 3.5 = int divide 7/2 = 3, then + 1
//  n = 5 | r = 3 | d -2
//  n = 3 | r = 2 | d -1
//  n = 1 | r = 1 | d 0
