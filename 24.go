package main

import (
  "fmt"
)

func main(){
  fac := []int{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800} // factorials of 0 to 10
  n:= 0
  fmt.Println("To find the nth lexograpic permutation of 0123456789 enter n:")
  fmt.Scanln(&n)
  fmt.Printf("n is %d\n", n)


 if n > 0 && n <= fac[2] {
   fmt.Println("two digits must change")
 }
 if n > fac[2] && n <= fac[3] {
   fmt.Println("three digits must change")
 }
 if n > fac[3] && n <= fac[4] {
   fmt.Println("four digits must change")
 }

//number of permutations for 0123456789
// 10 digits which can be arranged
//  so there are 10! permutations
// 10! == 3628800 total permutations

//try to count up from the smallest to the largest value using each digit once

// *permutations for last 2 digits*  // 2*1 = 2 | (2!) = 2 perms where 2 digits change

//0123456789 - index 1 - changed 2
//0123456798 - index 2 - changed 2

// there are 1! that start with 8
// there are 1! that start with 9

// *permutations for last 3 digits* // 3*2*1 = 6 |  (3!) - (2!) = 4 perms where 3 digits change

//0123456789 - index 1 - changed 2  - starts with 7 - also in perms for last 2
//0123456798 - index 2 - changed 2  - starts with 7 - also in perms for last 2
//0123456879 - index 3 - changed 3  - starts with 8
//0123456897 - index 4 - changed 3 -  starts with 8
//0123456978 - index 5 - changed 3 -  starts with 9
//0123456987 - index 6 - changed 3 -  starts with 9

// there are 2! that start with 7
// there are 2! that start with 8
// there are 2! that start with 9

// *permutations for last 4 digits*    // 4*3*2*1 = 24   | (4!) - (3!) = 18 perms where 4 digits change

//0123456789 - index 1 - changed 2  - starts with 6 - also in perms for last 2
//0123456798 - index 2 - changed 2  - starts with 6 - also in perms for last 2
//0123456879 - index 3 - changed 3  - starts with 6 - also in perms for last 3
//0123456897 - index 4 - changed 3  - starts with 6 - also in perms for last 3
//0123456978 - index 5 - changed 3  - starts with 6 - also in perms for last 3
//0123456987 - index 6 - changed 3  - starts with 6 - also in perms for last 3
//0123457689 - index 7 - changed 4  - starts with 7
//0123457698 - index 8 - changed 4  - starts with 7
//0123457869 - index 9 - changed 4  - starts with 7
//0123457896 - index 10 - changed 4 - starts with 7
//0123457968 - index 11 - changed 4 - starts with 7
//0123457986 - index 12 - changed 4 - starts with 7
//0123458679 - index 13 - changed 4 - starts with 8
//0123458697 - index 14 - changed 4 - starts with 8
//0123458769 - index 15 - changed 4 - starts with 8
//0123458796 - index 16 - changed 4 - starts with 8
//0123458967 - index 17 - changed 4 - starts with 8
//0123458976 - index 18 - changed 4 - starts with 8
//0123459678 - index 19 - changed 4 - starts with 9
//0123459687 - index 20 - changed 4 - starts with 9
//0123459768 - index 21 - changed 4 - starts with 9
//0123459786 - index 22 - changed 4 - starts with 9
//0123459867 - index 23 - changed 4 - starts with 9
//0123459876 - index 24 - changed 4 - starts with 9

// there are 3! that start with 6
// there are 3! that start with 7
// there are 3! that start with 8
// there are 3! that start with 9

// to find out how many digits need to change to reach n

//   0 < n <= 2             // 2 digits must change
//   2 < n <= 6             // 3 digits must change
//   6 < n <= 24            // 4 digits must change
//   24 < n <= 120          // 5 digits must change
//   120 < n <= 720         // 6 digits must change
//   720 < n <= 5040        // 7 digits must change
//   5040 < n <= 40320      // 8 digits must change
//   40320 < n <= 362880    // 9 digits must change
//   362880 < n <= 3628800  // 10 digits must change

// so 1,000,000 < 10*9*8*7*6*5*4*3*2*1
// 1,000,000 < 3,628,800
//  and 1,000,000 > 9*8*7*6*5*4*3*2*1
//   1,000,000 > 362,880
// this means the 1,000,000th lexographic permutation needs to change the location of all 10 digits in 0123456789

// p(2) + p(3) + p(4) + p(5) + p(6) + p(7) + p(8) + p(9) + p(10) = 3,628,800

// (2!) = 2 (number of perms that change 2 digits )
// (3!) - (2!) = 6 - 2 = 4 (number of perms that change 3 digits)
// (4!) - (3!) = 24 - 6 = 18 (number of perms that change 4 digits)
// (5!) - (4!) = 120 - 24 = 96 (number of perms that change 5 digits)
// (6!) - (5!) = 720 - 120 = 600 (number of perms that change 6 digits)
// (7!) - (6!) = 5040 - 720 = 4320 (number of perms that change 7 digits)
// (8!) - (7!) = 40320 - 5040 = 35280 (number of perms that change 8 digits)
// (9!) - (8!) = 362880 - 40320 = 322560 (number of perms that change 9 digits)
// (10!) - (9!) = 3628800 - 362880 = 3265920 (number of perms that change 10 digits)

// there are 3265920 perms that change 10 digits.
//  9! will start with 0  | 0 to (9!)             = indexes 0       -  362880
//  9! will start with 1  | (9! + 1) to 2(9!)     = indexes 362881  -  725760
//  9! will start with 2  | (2(9!) + 1) to 3(9!)  = indexes 725761  - 1088640
//  9! will start with 3  | (3(9!) + 1) to 4(9!)  = indexes 1088641 - 1451520
//  9! will start with 4  | (4(9!) + 1) to 5(9!)  = indexes 1451521 - 1814400
//  9! will start with 5  | (5(9!) + 1) to 6(9!)  = indexes 1814401 - 2177280
//  9! will start with 6  | (6(9!) + 1) to 7(9!)  = indexes 2177281 - 2540160
//  9! will start with 7  | (7(9!) + 1) to 8(9!)  = indexes 2540161 - 2903040
//  9! will start with 8  | (8(9!) + 1) to 9(9!)  = indexes 2903041 - 3265920
//  9! will start with 9  | (9(9!) + 1) to 10(9!) = indexes 3265921 - 3628800


// so this means the 1,000,000th will start with a 2
// next digit only has 9 possiblities so.. 8! = 40320
// 8! will start with a 0 | (2(9!) + 0) to (2(9!) + (8!))          = indexes 725760 to 766080
// 8! will start with a 1 | (2(9!) + 8! + 1) to (2(9!) + 2(8!))    = indexes 766081 to 806400
// 8! will start with a 3 | (2(9!) + 2(8!) + 1) to (2(9!) + 3(8!)) = indexes 806401 to 846720
// 8! will start with a 4 | (2(9!) + 3(8!) + 1) to (2(9!) + 4(8!)) = indexes 846721 to 887040
// 8! will start with a 5 | (2(9!) + 4(8!) + 1) to (2(9!) + 5(8!)) = indexes 887041 to 927360
// 8! will start with a 6 | (2(9!) + 5(8!) + 1) to (2(9!) + 6(8!)) = indexes 927361 to 967680
// 8! will start with a 7 | (2(9!) + 6(8!) + 1) to (2(9!) + 7(8!)) = indexes 967681 to 1008000
// 8! will start with a 8 | (2(9!) + 7(8!) + 1) to (2(9!) + 8(8!)) = indexes 1008001 to 1048320
// 8! will start with a 9 | (2(9!) + 8(8!) + 1) to (2(9!) + 9(8!)) = indexes 1048321 to 1088640

// this means that the second digit of the 1,000,000th lexographic permutation is a 7
// repeat the above step for 7! then add 2(9!) and  7(8!) to get the indexes
// can possibly use case statement to check which range nth lexograpic permutation falls within

}
