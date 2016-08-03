package main

import (
	"fmt"
)

func main() {
	for x := 900; x < 1000; x++ {
		for y := 900; y < 1000; y++ {
			xy := x * y
			fmt.Printf("xy is %d\n", xy)
			px := isPalindrome(xy)
			if px == true {
				fmt.Printf("%d is palindrome. Its products are %d and %d\n", xy, x, y)
			}
		}
	}

}

func isPalindrome(myInt int) bool {
	// get the digits and store them in a slice
	origInt := myInt
	digits := []int{}
	for myInt >= 1 {

		aDigit := myInt % 10
		digits = append(digits, aDigit)

		myInt = myInt - aDigit
		myInt = myInt / 10
	}

	// if the number being checked ends in 0 it can't be palindrome
	// because we are not checking any numbers that start with 0
	if digits[0] == 0 {
		return false
	}
	// reverse the digits and store in a new variable
	reverse := 0
	for i, _ := range digits {
		reverse = reverse * 10
		reverse = reverse + digits[i]
	}
	// compare the original digits to the reversed digits
	if reverse == origInt {
		return true
	}
	return false
}
