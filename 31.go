package main

import (
  "fmt"

)


func main(){

// using only 1p coins, the number of ways to pay 1p - 10p
//  0  1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//  1  1   1   1   1   1   1   1   1   1   1

// values array starts as
//  1  0  0   0   0   0   0   0   0   0   0
// starting at index[1] add  value at index[1 - 1]
//  1  1   1  1   1   1   1   1   1   1   1

// using only 1p and 2p coins the number of ways to pay 1p - 5p
//  0  1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//  1  1   2   2   3   3   4   4   5   5   6

// values array starts as
//  1  1   1   1   1   1   1   1   1  1   1
//  starting at index[2] add value at index [ -2]
//  1  1   2   2   3   3   4   4  5   5   6

// using only 1p, 2p and 5p coins the number of ways to pay 1p - 10p
//  0  1p, 2p, 3p, 4p, 5p, 6p, 7p, 8p, 9p, 10p
//  1   1   2   2   3   4   5   6   7   8   10

//values array starts as
//  1   1   2   2   3   3   4   4   5   5   6
// start at index[5] and add value at index [-5]
//  +0 +0  +0  +0  +0  +1  +1  +2  +3   +3  +4


// difference in number of ways between using 1p coins and using 1p and 2p coins
//  0  0  +1  +1  +2  +2  +3  +3  +4  +4  +5

// difference between using 1p and 2p and using 1p, 2p and 5p
//  0  0   0  0   0  +1  +1  +2  +2  +3  +4

// observations:
// for 1p coins, there is only 1 way to make each value
// when 2p coins are included, the totals up to but not including 2p do not change
// when 5p coins are included, the totals up to but not including 5p do not change

// how many ways can £2 be made from all these coins
amount := 200
coins := []int{1,2,5,10,20,50,100,200}
values := make([]int, amount+1) //  make a slice of size (amount + 1)
values[0] = 1 // set the first value in the combos slice to 1

// if amount is 10
// then values starts as {1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

    for i:=0; i < len(coins); i++ { // for each type of coin
    //  fmt.Println(coins[i]) // 1, 2, 5, 10, 50, 100, 200
       for j:= coins[i]; j <= amount; j++{ // for each value from current coin value to the total value
         values[j] += values[j - coins[i]] // add to the value in the values array, the value at index - coin value being checked
         }
    }
    fmt.Println(values[len(values)-1]) // the total number of ways the value can be made from the coins is the last value in the values array

}

// coins
