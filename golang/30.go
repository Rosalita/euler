package main

import (
  "fmt"
//  "strconv"

)

func main(){

// loop through numbers - ignore single digit numbers as these are not a sum
// take the first digit and find it to the power of 5
// take the second digit and find it to the power 5   -
// potentially write a function for power of 5 as will be doing this a lot -
// actually will only ever need to know the power of 5 for digits 0 - 9
// could probably just store 10 answers in a map so they dont have to be calculated each time
// add those two Values
// check if the sum is greater than the original number - if it is, move on to the next value
// if its not greater, see if there is another digit, if there isn't then this number is one we're searching for
// if there is another digit, take the next digit   find its power of 5 and add this
// repeat this loop until a large amount of numbers have been checked

powers5 := map[int]int{
 0 : 0,
 1 : 1,
 2 : 32,
 3 : 243,
 4 : 1024,
 5 : 3125,
 6 : 7776,
 7 : 16807,
 8 : 32768,
 9 : 59049,
}

limit := 9999999999999


  for i:=10; i <= limit; i++{
    slicedInt := makeDigitSlice(i)
  //  fmt.Println(slicedInt)
   total := 0
    for _, digit := range slicedInt{
      total += (powers5[digit])
    }

     if total > i || total < i {
       continue
     }

     fmt.Printf("%d can be written as the sum of 5th powers\n", i)

   }

}


func makeDigitSlice(x int) []int{

  digits := []int{} // create a blank slice of ints
  for x > 0 { // make sure the value being sliced is greater than 0
    nextDigit := x % 10  // the remainder of x divided by 10
    digits = append(digits, nextDigit) // add this to the slice
    x -= nextDigit // subtract it from x
    x /= 10 // then divide x by 10
  }

  // the slice created is backwards so swap it around
  for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
    digits[i], digits[j] = digits[j], digits[i]
  }
  return digits
}
