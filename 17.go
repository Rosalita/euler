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

	// start counting and add the names of the numbers to a string
	// use a buffer to add them because its much more efficient than using +=
	// create this buffer using the Buffer type from package "bytes"
	var buffer bytes.Buffer
	var numberNames string

	for i := 1; i <= 99; i++ { //start counting

		if i < 21 { // for the first 20 numbers:
			buffer.WriteString(dict[i]) //write name of each number into the buffer
			continue                    // then continue to next number
		}
		if (i%10 == 0) && (i < 100) { // for numbers divisible by 10 that are less than 100
			buffer.WriteString(dict[i]) // its name will be in the dictionary so just write it to buffer
			continue                    // then continue to the next number
		}

		if i < 100 { // for numbers less than 100 which havent been written yet
			tens := i / 10
			switch tens {
			case 2:
				buffer.WriteString(dict[20])
				buffer.WriteString(dict[i-20])
			case 3:
				buffer.WriteString(dict[30])
				buffer.WriteString(dict[i-30])
			case 4:
				buffer.WriteString(dict[40])
				buffer.WriteString(dict[i-40])
			case 5:
				buffer.WriteString(dict[50])
				buffer.WriteString(dict[i-50])
			case 6:
				buffer.WriteString(dict[60])
				buffer.WriteString(dict[i-60])
			case 7:
				buffer.WriteString(dict[70])
				buffer.WriteString(dict[i-70])
			case 8:
				buffer.WriteString(dict[80])
				buffer.WriteString(dict[i-80])
			case 9:
				buffer.WriteString(dict[90])
				buffer.WriteString(dict[i-90])
			}

			continue //then continue to the next number
		}

	}
	numberNames = buffer.String() // take the value out of the buffer and put in a variable
	fmt.Printf("numberNames is %s\n", numberNames)
	fmt.Printf("the number of letters used is %d\n", len(numberNames))
}
