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

	for i := 14; i > 0; i-- {

		largestNeswSum, firstNesw, secondNesw, largestNwseSum, firstNwse, secondNwse, neswpos1, neswpos2, nwsepos1, nwsepos2 := getBiggestSums(i, triangleRows)

		fmt.Printf("for row %d and %d:\n", i, i-1)
		fmt.Printf("largest sw>ne sum is %d which is %d + %d at positions %d and %d\n", largestNeswSum, firstNesw, secondNesw, neswpos1, neswpos2)
		fmt.Printf("largest se>nw sum is %d which is %d + %d at positions %d and %d\n", largestNwseSum, firstNwse, secondNwse, nwsepos1, nwsepos2)
	}

	journey := getBiggestNumberPath(9, 14, triangleRows)
	fmt.Printf("journey is %s", journey)
}

func getBiggestNumberPath(startPoint, startRow int, triangleRows []string) (journey []string) {
	nextPoint := 0
	fmt.Printf("startpoint is %d\n", startPoint)
	rowSlice := strings.Split(triangleRows[startRow], " ")
	fmt.Printf("value at start point is %s\n", rowSlice[startPoint])
	// add the starting value to the journey
	journey = append(journey, rowSlice[startPoint])

	//convert this value to an int and add it to the total
	intVal, err := strconv.ParseInt(rowSlice[startPoint], 10, 64)
	if err != nil {
		panic(err)
	}
	var total int64
	total = total + intVal

	prevRowSlice := strings.Split(triangleRows[startRow-1], " ")
	// check which of the two values can move to is largest
	fmt.Printf("can move to %s or %s\n", prevRowSlice[startPoint-1], prevRowSlice[startPoint])
	// convert these string values to int64
	option1, err := strconv.ParseInt(prevRowSlice[startPoint-1], 10, 64)
	if err != nil {
		panic(err)
	}
	option2, err := strconv.ParseInt(prevRowSlice[startPoint], 10, 64)
	if err != nil {
		panic(err)
	}
	if option1 > option2 {
		journey = append(journey, prevRowSlice[startPoint-1])
		nextPoint = startPoint - 1
	} else {
		journey = append(journey, prevRowSlice[startPoint])
		nextPoint = startPoint
	}

	// check the next point - possibly start looping back through points here
	fmt.Printf("next point is %d\n", nextPoint)
	nextSlice := strings.Split(triangleRows[startRow-1], " ")
	fmt.Printf("value at next point is %s\n", nextSlice[nextPoint])
	nextPrevRowSlice := strings.Split(triangleRows[startRow-2], " ")
	// check which of the two values can move to is largest
	fmt.Printf("can move to %s or %s\n", nextPrevRowSlice[nextPoint-1], nextPrevRowSlice[nextPoint])
	// convert these string values to int64
	option1, err = strconv.ParseInt(nextPrevRowSlice[nextPoint-1], 10, 64)
	if err != nil {
		panic(err)
	}
	option2, err = strconv.ParseInt(nextPrevRowSlice[nextPoint], 10, 64)
	if err != nil {
		panic(err)
	}
	if option1 > option2 {
		journey = append(journey, nextPrevRowSlice[nextPoint-1])
		nextPoint = nextPoint - 1
	} else {
		journey = append(journey, nextPrevRowSlice[nextPoint])
	}

	return journey
}

func getBiggestSums(startRow int, triangleRows []string) (neswSum, firstNesw, secondNesw, nwseSum, firstNwse, secondNwse int64, neswpos1, neswpos2, nwsepos1, nwsepos2 int) {

	//make rowSlice which is a slice containing all the numbers on the startRow of triangleRows
	rowSlice := strings.Split(triangleRows[startRow], " ")
	//make prevRowSlice which is a slice of all the numbers in the previous row of triangleRows
	prevRowSlice := strings.Split(triangleRows[startRow-1], " ")
	// add a 0 on to the end of the previous row slice so both slices have the same number of values
	prevRowSlice = append(prevRowSlice, "0")

	neswSumSlice := make([]int64, 15) // used to store the sum of north east to south west direction number pairs
	nwseSumSlice := make([]int64, 15) // used to store the sum of north west to south east direction number pairs

	// calculate the sums of north east to south west number pairs and store them as int64 in neswSum[]
	for i, _ := range rowSlice {

		rowInt, err := strconv.ParseInt(rowSlice[i], 10, 64) // convert the numeric string to an int64 and store in rowInt
		if err != nil {
			panic(err)
		}

		prevRowInt, err := strconv.ParseInt(prevRowSlice[i], 10, 64) // convert the numeric string to an int64 and store in prevRowInt
		if err != nil {
			panic(err)
		}

		total := rowInt + prevRowInt // add the two ints together
		neswSumSlice[i] = total      // and store the sum in neswSumSlice
	}

	// calculate the sums of north west to south east number pairs and store them as int64 in nwseSum[]
	for i, _ := range rowSlice {
		var prevRowInt int64
		var err error

		rowInt, err := strconv.ParseInt(rowSlice[i], 10, 64) // convert the numeric string to an int64 and store in rowInt

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
		nwseSumSlice[i] = total
	}

	var largestNesw int64 = 0
	var largestNwse int64 = 0
	sfirstNesw := " "
	ssecondNesw := " "
	sfirstNwse := " "
	ssecondNwse := " "
	neswPos1, neswPos2, nwsePos1, nwsePos2 := 0, 0, 0, 0

	//find the largest nesw value in the neswSumSlice
	for i, v := range neswSumSlice {
		if i == 0 {
			largestNesw = v          // remember the largest value
			sfirstNesw = rowSlice[i] //and the numbers that were added to make that value
			neswPos1 = i
			ssecondNesw = prevRowSlice[i] // these numbers are actually strings
			neswPos2 = i
		} else {
			if v > largestNesw {
				largestNesw = v
				sfirstNesw = rowSlice[i]
				neswPos1 = i
				ssecondNesw = prevRowSlice[i]
				neswPos2 = i
			}
		}
	}

	//find the largest nwse value in the nwseSumSlice
	for i, v := range nwseSumSlice {
		if i == 0 {
			largestNwse = v // remember the largest value
			sfirstNwse = rowSlice[i]
			nwsePos1 = i
			ssecondNwse = "0"
			nwsePos2 = i
		} else {
			if v > largestNwse {
				largestNwse = v
				sfirstNwse = rowSlice[i]
				nwsePos1 = i
				ssecondNwse = prevRowSlice[i-1]
				nwsePos2 = i - 1

			}
		}
	}

	// convert all the numeric strings to int64 before returning them
	ifirstNesw, err := strconv.ParseInt(sfirstNesw, 10, 64)
	if err != nil {
		panic(err)
	}
	isecondNesw, err := strconv.ParseInt(ssecondNesw, 10, 64)
	if err != nil {
		panic(err)
	}
	ifirstNwse, err := strconv.ParseInt(sfirstNwse, 10, 64)
	if err != nil {
		panic(err)
	}
	isecondNwse, err := strconv.ParseInt(ssecondNwse, 10, 64)
	if err != nil {
		panic(err)
	}
	// return largest sum northeast to southwest and the two values that make this sum
	// also return the largest sum northwest to southeast and the two values that make this sum
	// return the position 0 - startRow of the numbers which when added make the largest sum
	return largestNesw, ifirstNesw, isecondNesw, largestNwse, ifirstNwse, isecondNwse, neswPos1, neswPos2, nwsePos1, nwsePos2
}
