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

	rawTriangleData, err = ioutil.ReadFile("triangle1.txt")

	if err != nil {
		panic(err)
	}
	stringTriangleData := string(rawTriangleData)
	triangleRows := strings.Split(stringTriangleData, "\r\n")

	fmt.Printf("triangle row 14: %v\n", triangleRows[14])
		fmt.Printf("triangle row 13: %v\n", triangleRows[13])
	//var rowSlice []string
	rowSlice := strings.Split(triangleRows[14], " ")
	prevRowSlice := strings.Split(triangleRows[13], " ")
	// add a 0 on to the end of the slice for prev row so both rows have the same number of values
 prevRowSlice = append(prevRowSlice, "0")

	neswSum := make([]int64, 15) // used to store the sum of north east to south west direction number pairs
	nwseSum := make([]int64, 15) // used to store the sum of north west to south east direction number pairs

// calculate the sums of north east to south west number pairs and store them as int64 in neswSum[]
	for i, _ := range rowSlice {
//		fmt.Printf("rowSlice %d is %v\n", i, v)
		rowInt, err := strconv.ParseInt(rowSlice[i], 10, 64)
		if err != nil {
			panic(err)
		}

		prevRowInt, err := strconv.ParseInt(prevRowSlice[i], 10, 64)
		if err != nil {
			panic(err)
		}
//		fmt.Printf("prevRowSlice %d is %v\n", i, prevRowSlice[i])
		total := rowInt + prevRowInt
		neswSum[i] = total
	}

// calculate the sums of north west to south east number pairs and store them as int64 in nwseSum[]
for i, _ := range rowSlice {
	  var prevRowInt int64
		var err error

		rowInt, err := strconv.ParseInt(rowSlice[i], 10, 64)

		if err != nil {
			panic(err)
		}

		// offset the previous row to start with 0 so that direction of addition changes from nesw to nwse
		if i == 0 {
			prevRowInt = 0
		} else {
			prevRowInt, err = strconv.ParseInt(prevRowSlice[i-1], 10, 64)
			if err != nil {
				panic(err)
			}
		}
   total := rowInt + prevRowInt
	 nwseSum[i] = total
}

var largestNesw int64 = 0
var largestNwse int64 = 0
firstNesw := " "
secondNesw := " "
firstNwse := " "
secondNwse := " "

//find the largest nesw value
for i, v := range neswSum {
	if i == 0{
		largestNesw = v
		firstNesw = rowSlice[i]
		secondNesw = prevRowSlice[i]
	} else {
		if v > largestNesw {
			largestNesw = v
			firstNesw = rowSlice[i]
			secondNesw = prevRowSlice[i]
		}
	}
}

//find the largest nwse value
for i, v := range nwseSum {
	if i == 0{
		largestNwse = v
		firstNwse = rowSlice[i]
		secondNwse = "0"
	} else {
		if v > largestNwse {
			largestNwse = v
			firstNwse = rowSlice[i]
			secondNwse = prevRowSlice[i-1]

		}
	}
}
	fmt.Printf("largest sum of nesw is %d \n", largestNesw)
	fmt.Printf("this is the sum of %s and %s\n", firstNesw, secondNesw)
	fmt.Printf("largest sum of nwse is %d \n", largestNwse)
	fmt.Printf("this is the sum of %s and %s\n", firstNwse, secondNwse)

}
