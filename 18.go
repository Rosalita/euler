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

	for i := 0; i < len(triangleRows)-1; i++ { // for each possible starting point in the final row ( 0 - 14)

		journey := getBiggestNumberPath(i, len(triangleRows)-1, triangleRows)
		fmt.Printf("journey is %s\n", journey)

		var totalJourney int64
		for i := 0; i < len(journey); i++ {
			intVal, err := strconv.ParseInt(journey[i], 10, 64)
			if err != nil {
				panic(err)
			}
			totalJourney += intVal
		}
		fmt.Printf("Total Journey is %d\n", totalJourney)

	}
}

func getBiggestNumberPath(startPoint, startRow int, triangleRows []string) (journey []string) {
	// take a slice of values from the starting row
	rowSlice := strings.Split(triangleRows[startRow], " ")

	// add the starting point value in the starting row to the journey
	journey = append(journey, rowSlice[startPoint])

	for i := 14; i > 1; i-- { // loop from the starting row to the the penultimate row

		//get the slice of the current row
		rowSlice = strings.Split(triangleRows[i], " ")
		// get the slice of the previous row
		prevRowSlice := strings.Split(triangleRows[i-1], " ")

		if startPoint == 0 {
			// do nothing to startPoint
			// there is only 1 move available
			journey = append(journey, prevRowSlice[startPoint])
		} else if startPoint == len(rowSlice)-1 { //if the starting point is at the end of the rowSlice
			//then only one move is available

			// can only move to one possible open
			// add this value to journey
			journey = append(journey, prevRowSlice[startPoint-1])
			if startPoint != 0 {
				startPoint -= 1 // dont let start point fall below 0!
			}
			// start point must be decremented to stay in range of the slices
		} else { // otherwise there are two moves available

			//convert the two string number for both these options to int64
			option1, err := strconv.ParseInt(prevRowSlice[startPoint-1], 10, 64)
			if err != nil {
				panic(err)
			}
			option2, err := strconv.ParseInt(prevRowSlice[startPoint], 10, 64)
			if err != nil {
				panic(err)
			}
			// compare the two choices
			if option1 == option2 {
				//add one of the options to the journey it doesnt matter which
				journey = append(journey, prevRowSlice[startPoint])
				fmt.Println("unable to decide which direction both values same")
				fmt.Println("so not changing direction")
			}

			if option1 > option2 { // if option 1 is the largest number
				journey = append(journey, prevRowSlice[startPoint-1]) // add it to the journey
				if startPoint != 0 {                                  // bring the starting for the next loop in by 1
					startPoint -= 1 // but dont let start point fall below 0!
				}
			} else {
				journey = append(journey, prevRowSlice[startPoint]) // add the other number to the journey
			}

		}
		if i == 2 { // add the number at the very top of the triangle to the journey
			rowSlice = strings.Split(triangleRows[0], " ")
			journey = append(journey, rowSlice[0])
		}
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
