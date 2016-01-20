package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("13.txt")
	if err != nil {
		panic(err)
	}

	stringData := string(data)
	slicesOfRows := strings.Split(stringData, "\n")

	// add the last digits of all the numbers
	sum := addNthDigitFromAllRows(slicesOfRows, 49)
	fmt.Printf("sum of last digit in all rows is %d\n", sum)
	tens := 0
	second := 0
	//third := 0

	lastdigit := int(sum % 10)        // get the last digit of the sum
	sdigit := strconv.Itoa(lastdigit) // convert it to string

	// store the answer as a slice of strings
	var answer []string
	answer = append(answer, sdigit)

	// if the sum all the numbers is greater than 10
	// store the second digit of the sum to be used in next calc
	if sum > 10 { //422
		//remove the last digit by subtracting it and dividing by 10
		tens = (int(sum) - lastdigit) / 10
		fmt.Printf("tens %d\n", tens)
		second = tens % 10
		fmt.Printf("second digit of sum is %d\n", second)
	}
	// if the sum of all the numbers is greater than 100
	// store the third digit of the sum to be used in next calc

	if sum > 100 { //422
		third := (int(tens) - second) / 10
		fmt.Printf("third digit of sum is %d\n", third)
	}

	fmt.Println(string(answer[0]))
}

func addNthDigitFromAllRows(rows []string, n int) int64 {
	var sum int64
	for i := 0; i < 100; i++ {
		number := getNthDigit(rows[i], n)
		sum = sum + number

	}
	return sum
}

func getNthDigit(myString string, n int) int64 {
	digit := string(myString[n])
	intdigit, ok := strconv.ParseInt(digit, 10, 0)
	if ok != nil {
		panic(ok)
	}
	return intdigit
}
