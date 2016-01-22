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


  var sums []string
  modifier := 1000000000
	total:= 0
	for i:= 0; i < 10; i++ {
		columnsum := int(addNthDigitFromAllRows(slicesOfRows, i))
		columnsum = columnsum * modifier
		total = total + columnsum
		modifier = modifier / 10
		stringsum := strconv.Itoa(columnsum)
	// append it to slice of sums
		sums = append(sums, stringsum)
	}

	for i:= 0; i < 10; i++ {
		fmt.Println(slicesOfRows[i])
		fmt.Println(sums[i])
	}
 fmt.Printf("total is %d", total)

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
