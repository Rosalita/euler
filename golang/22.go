package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
  "time"
)

func main() {
  start := time.Now()
	data, err := ioutil.ReadFile("names.txt") // read in the data from text file
  if err != nil { // if there is an error, handle it
		panic(err)
	}
	strData := string(data)                   // convert the []uint8 bytes of data to a string
	strData = strData[1 : len(strData)-1]     // remove first and last characters
	names := strings.Split(strData, "\",\"")  // split the data on "," into []string
	sort.Strings(names)                       // sort the names into alphabetical order
	totalScore := 0 // initialise a total
	for i, v := range names { // for each name
		scoreLetters := scoreLetters(v) //score the letter
		nameScore := (i + 1) * scoreLetters // multiply letter score by position
		totalScore += nameScore // add name score to total
	}
	fmt.Println(totalScore)
  fmt.Println("Elapsed time : ", time.Since(start))
}

func scoreLetters(name string) int {
	total := 0
	for _, v := range name {
		total += getLetterValue(string(v))
	}
	return total
}

func getLetterValue(letter string) int {
letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
value := strings.Index(letters, letter)
return value + 1
}
