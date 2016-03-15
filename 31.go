package main

import (
  "fmt"

)


func main(){

// using only 1p coins, the number of ways to pay 1p - 10p
//   1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//   1   1   1   1   1   1   1   1   1   1

// using only 1p and 2p coins the number of ways to pay 1p - 5p
//   1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//   1   2   2   3   3   4   4   5   5   6

// using only 1p, 2p and 5p coins the number of ways to pay 1p - 10p
//   1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//   1   2   2   3   4   5   6   7   8   10


// difference in number of ways between using 1p coins and using 1p and 2p coins
//   0  +1  +1  +2  +2  +3  +3  +4  +4  +5

// difference between using 1p and 2p and using 1p, 2p and 5p
//   0   0  0   0  +1  +1  +2  +2  +3  +4

// how many ways can £2 be made from all these coins
//amount := 200
coins := []int{1,2,5,10,20,50,100,200}
    for i:=0; i < len(coins); i++ { // for each type of coin
      fmt.Println(coins[i])
    }
}
