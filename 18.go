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

	fmt.Printf("triangle row 14 %v\n", triangleRows[14])
	//var rowSlice []string
	rowSlice := strings.Split(triangleRows[14], " ")
	prevRowSlice := strings.Split(triangleRows[13], " ")
	sumOfSlices := make([]int64, 15)

	for i, v := range rowSlice {
		if i == 14 {
			break
		}
		fmt.Printf("rowSlice %d is %v\n", i, v)
		rowInt, err := strconv.ParseInt(rowSlice[i], 10, 64)
		if err != nil {
			panic(err)
		}

		prevRowInt, err := strconv.ParseInt(prevRowSlice[i], 10, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("prevRowSlice %d is %v\n", i, prevRowSlice[i])
		total := rowInt + prevRowInt
		sumOfSlices[0] = total
	}

}
