package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	factorial := new(big.Int)          // make a new big.Int
	factorial.MulRange(1, 100)         // The Go package math.big has method MulRange which when called with first param 1, it returns the second param as its factorial value
	strFactorial := factorial.String() // a big.Int can be converted to a string using it's String() method

	var total int64 = 0 // intialise a variable to count

	for _, v := range strFactorial { // for each character in the string strFactorial
		sv := string(v)                       // convert the value which is type byte to type string
		iv, _ := strconv.ParseInt(sv, 10, 64) // parse the string value into an int64
		total = total + iv                    // add each value to the total
	}
	fmt.Println(total)
}
