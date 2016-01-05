package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	var x, lrgfactor int
	fmt.Println("to find largest prime factor for x, enter a value for x:")
	fmt.Scanln(&x)
	// remember []uint8 is the same as []byte
	var primes []uint8
	var err error
	primes, err = ioutil.ReadFile("primes-to-100k.txt")
	    if err != nil {
		    panic(err)
	    }
	primestr := string(primes) // converts []uint8 to a string
	primenumbers := strings.Split(primestr, "\n") // slices the string  into []string

	for i, _ := range primenumbers {
		num := primenumbers[i] // num is now a string containing "<value>\r"
		trimmedPrime := strings.TrimSpace(num) // removes the whitespace "\r" from the end of the string
		int64Prime, err := strconv.ParseInt(trimmedPrime, 10, 0) //convert the string to an int64, ParseInt always returns int64
		    if err != nil {
					panic(err)
				}
		intPrime := int(int64Prime) // convert int64 to an int

		//finally can check if each prime number is a factor of x
		if x % intPrime == 0 {
				fmt.Printf("%d is a factor of %d\n", intPrime, x)
				lrgfactor = intPrime
		}
	}
	fmt.Printf("The largest prime factor of %v is %v", x, lrgfactor)
}
