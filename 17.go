package main

import (
	"bytes"
	"fmt"
)

func main() {
	var dict map[int]string // create a map which is a dictionary
	// teach it all the words it needs to generate numbers upto 1000
	dict = map[int]string{1: "one", 2: "two", 3: "three",
		4: "four", 5: "five", 6: "six",
		7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve",
		13: "thirteen", 14: "fourteen", 15: "fifteen",
		16: "sixteen", 17: "seventeen", 18: "eighteen",
		19: "nineteen", 20: "twenty", 30: "thirty",
		40: "forty", 50: "fifty", 60: "sixty",
		70: "seventy", 80: "eighty", 90: "ninety",
		100: "hundred", 1000: "thousand"}

	//  fmt.Printf(" %s %s %s %s %s",dict[1],dict[2],dict[3],dict[4],dict[5])

	// start counting and add the names of the numbers to a big long string
	// use a buffer to add them because its much more efficient than using +=
	// create this buffer using the Buffer type from package "bytes"
	var buffer bytes.Buffer
	var numberNames string

	for i := 1; i <= 1000; i++ { //start counting

		if i == 1000 {
			buffer.WriteString(dict[1])
			buffer.WriteString(dict[i])
		}

		if i < 1000 {
			hundreds := i / 100
			tens := i/10 - (hundreds * 10)
			units := i - (tens * 10) - (hundreds * 100)
			tensandunits := (tens * 10) + units

			if hundreds > 0 {
				buffer.WriteString(dict[hundreds])
				buffer.WriteString(dict[100])
			}
			if (hundreds > 0) && (tens > 0) || (hundreds > 0) && (units > 0) {
				buffer.WriteString("and")
			}

			if tens > 0 {
				if (tensandunits > 10) && (tensandunits < 20) {
					buffer.WriteString(dict[tensandunits])
					continue
				} else {
					buffer.WriteString(dict[tens*10])
				}

			}
			if units > 0 {
				buffer.WriteString(dict[units])
			}
		}
	}
	numberNames = buffer.String() // take the value out of the buffer and put in a variable
	fmt.Printf("numberNames is %s\n", numberNames)
	fmt.Printf("the number of letters used is %d\n", len(numberNames))
}
