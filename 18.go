package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var rawTriangleData []byte
	var err error

	// read the triangle to be solved from a text file
	rawTriangleData, err = ioutil.ReadFile("triangle1.txt")
	if err != nil {
		panic(err)
	}

	//convert the raw bytes into string format
	stringTriangleData := string(rawTriangleData)

	//break the raw string down into a slice containing of rows of numbers
	triangleRows := strings.Split(stringTriangleData, "\r\n")

	answer := getBiggestSum(triangleRows)
	fmt.Println(answer)

}

func getBiggestSum(triangleRows []string) string {

	for i := len(triangleRows) - 2; i >= 0; i-- {
		updatingRow := triangleRows[i]
		comparingRow := triangleRows[i+1]

		//split the updating row into each number
		updatingNumbers := strings.Split(updatingRow, " ")
		//split the comparing row into each number
		comparingNumbers := strings.Split(comparingRow, " ")

		//loop through each number in the updating row so they can be updated
		for i := 0; i < len(updatingNumbers); i++ {
			// if the first number checked is the biggest of the two, dd it to the updatednumber
			if comparingNumbers[i] > comparingNumbers[i+1] {
				// convert both strings to int
				cnum, err := strconv.ParseInt(comparingNumbers[i], 10, 64)
				if err != nil {
					panic(err)
				}
				unum, err := strconv.ParseInt(updatingNumbers[i], 10, 64)
				if err != nil {
					panic(err)
				}
				sum := int(cnum + unum) // add them together
				s := strconv.Itoa(sum)  // convert back to string
				updatingNumbers[i] = s  // store new value in updatingNumbers slice

			}

			// if the second number checked is the biggest of the two, dd it to the updatednumber
			if comparingNumbers[i+1] > comparingNumbers[i] {
				// convert both strings to int
				cnum, err := strconv.ParseInt(comparingNumbers[i+1], 10, 64)
				if err != nil {
					panic(err)
				}
				unum, err := strconv.ParseInt(updatingNumbers[i], 10, 64)
				if err != nil {
					panic(err)
				}

				sum := int(cnum + unum) // add them together
				s := strconv.Itoa(sum)  // convert back to string
				updatingNumbers[i] = s  // store new value in updatingNumbers slice
			}

			// if both numbers checked are the same, add this number once to the updatednumber
			if comparingNumbers[i+1] == comparingNumbers[i] {
				// convert one of the strings checked to int
				cnum, err := strconv.ParseInt(comparingNumbers[i+1], 10, 64)
				if err != nil {
					panic(err)
				}
				// and convert the number the string is being added with
				unum, err := strconv.ParseInt(updatingNumbers[i], 10, 64)
				if err != nil {
					panic(err)
				}
				sum := int(cnum + unum) // add them together
				s := strconv.Itoa(sum)  // convert back to string
				updatingNumbers[i] = s  // store new value in updatingNumbers slice
			}

		}
		// join the updated numbers back together
		updatedRow := strings.Join(updatingNumbers, " ")
		// write the updatedNumbers string back into triangleRows
		triangleRows[i] = updatedRow

		if i == 0 {
			return triangleRows[0]
		}
	}
	return "fish"
}
