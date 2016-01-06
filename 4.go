package main

import (
  "fmt"
)

func main(){
//  s:= []int{}
  for x := 900; x < 1000; x++ {
     for y := 900; y < 1000; y++ {
       xy := x * y
       fmt.Printf("xy is %d\n", xy)
    //   s = append(s, xy)
       px:= isPalindrome(xy)
         if px == true {

         fmt.Printf("%d is palindrome. Its products are %d and %d\n",xy, x, y)
         }
       }
     }

}

func isPalindrome(myInt int) bool{
      // get the digits and store them in a slice
      origInt := myInt
      digits:= []int{}
      for myInt >= 1{

      aDigit := myInt % 10
      digits = append(digits, aDigit)

      myInt = myInt - aDigit
      myInt = myInt / 10
      }
      // reverse the digits and store in a new variable
      reverse := 0
      for i, _ := range digits{
        reverse = reverse * 10
        reverse = reverse + digits[i]
      }
     // compare the original digits to the reversed digits
    // fmt.Printf("reverse is %d origint is %d", reverse, origInt)
      if reverse == origInt {
        return true
      }
      return false
}
