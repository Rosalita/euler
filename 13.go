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

	// add the first digit in all the rows

	fmt.Printf("sum of first digit in all rows is %d", addNthDigitFromAllRows(slicesOfRows, 0))

}

func addNthDigitFromAllRows(rows []string, n int) int64 {
	var sum int64
	for i := 0; i < 100; i++ {
		number := getNthDigit(rows[i], 0)
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
