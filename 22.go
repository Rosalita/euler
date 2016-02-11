package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("names.txt") // read in the data from text file
	handleError(err)                          // if there is an error, handle it
	strData := string(data)                   // convert the []uint8 bytes of data to a string
	strData = strData[1 : len(strData)-1]     // remove first and last characters
	names := strings.Split(strData, "\",\"")  // split the data on "," into []string
	sort.Strings(names)                       // sort the names into alphabetical order
	totalScore := 0
	for i, v := range names {
		scoreLetters := scoreLetters(v)
		nameScore := (i + 1) * scoreLetters
		totalScore += nameScore
	}
	fmt.Println(totalScore)
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func scoreLetters(name string) int {
	total := 0
	for _, v := range name {
		total += getLetterValue(string(v))
	}
	return total
}

func getLetterValue(letter string) int {
	switch letter {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	case "D":
		return 4
	case "E":
		return 5
	case "F":
		return 6
	case "G":
		return 7
	case "H":
		return 8
	case "I":
		return 9
	case "J":
		return 10
	case "K":
		return 11
	case "L":
		return 12
	case "M":
		return 13
	case "N":
		return 14
	case "O":
		return 15
	case "P":
		return 16
	case "Q":
		return 17
	case "R":
		return 18
	case "S":
		return 19
	case "T":
		return 20
	case "U":
		return 21
	case "V":
		return 22
	case "W":
		return 23
	case "X":
		return 24
	case "Y":
		return 25
	case "Z":
		return 26
	}
	panic("unable to get a value for letter")
}
