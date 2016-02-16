package main

import (
  "fmt"
)

func main(){

//number of permutations for 0123456789
// 10 digits which can be arranged
//  so there are 10! permutations
// 10! == 3628800

//try to count up from the smallest to the largest value using each digit once

// *permutations for last 2 digits*  // 2*1 = 2 | p(2) = 2

//0123456789 - index 1 - changed 2
//0123456798 - index 2 - changed 2

// *permutations for last 3 digits* // 3*2*1 = 6 |  p(2) + p(3) = 6

//0123456789 - index 1 - changed 2  - also in perms for last 2
//0123456798 - index 2 - changed 2  - also in perms for last 2
//0123456879 - index 3 - changed 3
//0123456897 - index 4 - changed 3
//0123456978 - index 5 - changed 3
//0123456987 - index 6 - changed 3

// *permutations for last 4 digits*    // 4*3*2*1 = 24   | p(2) + p(3) + p(4) = 24

//0123456789 - index 1 - changed 2  - also in perms for last 2
//0123456798 - index 2 - changed 2  - also in perms for last 2
//0123456879 - index 3 - changed 3  - also in perms for last 3
//0123456897 - index 4 - changed 3  - also in perms for last 3
//0123456978 - index 5 - changed 3  - also in perms for last 3
//0123456987 - index 6 - changed 3  - also in perms for last 3
//0123457689 - index 7 - changed 4
//0123457698 - index 8 - changed 4
//0123457869 - index 9 - changed 4
//0123457896 - index 10 - changed 4
//0123457968 - index 11 - changed 4
//0123457986 - index 12 - changed 4
//0123458679 - index 13 - changed 4
//0123458697 - index 14 - changed 4
//0123458769 - index 15 - changed 4
//0123458796 - index 16 - changed 4
//0123458967 - index 17 - changed 4
//0123458976 - index 18 - changed 4
//0123459678 - index 19 - changed 4
//0123459687 - index 20 - changed 4
//0123459768 - index 21 - changed 4
//0123459786 - index 22 - changed 4
//0123459867 - index 23 - changed 4
//0123459876 - index 24 - changed 4

// so 1,000,000 < 10*9*8*7*6*5*4*3*2*1
// 1,000,000 < 3,628,800
//  and 1,000,000 > 9*8*7*6*5*4*3*2*1
//   1,000,000 > 362,880
// this means the 1,000,000th lexographic permutation needs to change the location of all 10 digits in 0123456789

// p(2) + p(3) + p(4) + p(5) + p(6) + p(7) + p(8) + p(9) + p(10) = 3,628,800 


}
