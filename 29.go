package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var start, end float64 = 2, 5
	var terms []float64

	for i := start; i <= end; i++ { // loop through all integers
		for j := start; j <= end; j++ { // loop through all powers
			pow := math.Pow(i, j)
			terms = append(terms, pow)
		}
	}
	sort.Float64s(terms) // sort the terms in numerical order
	deleted := 0
	finishflag := false
	for i, _ := range terms { // for each value in the slice of terms
		if finishflag == true { // if flag to finish looping is true, use continue to skip any remaining values
			continue
		}
		if i == 0 { // if first value, unable to compare it to previous value so skip it with continue
			continue
		}
		// if anything has been deleted from the original slice, deleted will be > 0
		// for each deletion the slice will be shorter so...
		for d := 1; d <= deleted; d++ { //for each time a value was deleted
			if i == len(terms)-1-d { // if terms[i] would be at the index of (last value) - (number of items deleted)
				finishflag = true // set finishflag to true
				continue          // and skip this value
			}
		}

		if terms[i] == terms[i-1] { // compare the value to the one before it and if its the same
			terms = append(terms[:i], terms[i+1:]...) // delete it by setting the slice to all values before and after the term to be removed
			deleted++ // increment a counter of deleted items
    }
	}
  	//fmt.Println(terms)
      count := 0
      for _, _ = range(terms){
        count ++
      }
      fmt.Printf("There are %d distinct terms\n", count)
}
