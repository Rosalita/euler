package main

import (
	"bytes"
	"fmt"
)

func main() {
	var dict map[int]string
	dict = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five"}
	//  fmt.Printf(" %s %s %s %s %s",dict[1],dict[2],dict[3],dict[4],dict[5])

	// add the names of numbers 1 to 5 to a string
	// use a buffer for this because its much more efficient than +=
	// the bytes package has a Buffer type, its zero value is an empty ready to use buffer
	var buffer bytes.Buffer
	var numberNames string

	for i := 1; i <= 5; i++ {
		buffer.WriteString(dict[i]) //write each item into the buffer
	}
	numberNames = buffer.String() // take the value out of the buffer
	fmt.Printf("numberNames is %s\n", numberNames)
	fmt.Printf("the number of letters used is %d\n", len(numberNames))
}
