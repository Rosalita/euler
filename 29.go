package main

import (
	"fmt"
  "math/big"
	"sort"
)

//create a new type which is a slice of pointers to big.Int types
type bigIntSlice []*big.Int

// add Len() Less and Swap methods to bigIntSlice so it fits sort.Interface and can be sorted
func(n bigIntSlice) Len() int{ return len(n) }
func(n bigIntSlice) Less(i,j int) bool{ return n[i].Int64() < n[j].Int64() }
func(n bigIntSlice) Swap(i,j int) { n[i], n[j] = n[j], n[i] }

func main() {
	var start, end int = 2, 5
	var bigTerms bigIntSlice // slice of pointers to Int type from math/big package

	for i := start; i <= end; i++ { // loop through all integers
		for j := start; j <= end; j++ { // loop through all powers
      big1, big2, bigAns:= big.NewInt(int64(i)), big.NewInt(int64(j)), big.NewInt(int64(1))
      bigAns.Exp(big1, big2, nil)
			bigTerms = append(bigTerms, bigAns)
		}
	}

  sort.Sort(bigIntSlice(bigTerms)) // sort the terms into numerical order
	fmt.Println(bigTerms)
	fmt.Printf("there are %d bigTerms before deletion\n", len(bigTerms))
	deleted := 0
	finishflag := false
	for i, _ := range bigTerms { // for each value in the slice of terms
		if finishflag == true { // if flag to finish looping is true, use continue to skip any remaining values
			continue
		}
		if i == 0 { // if first value, unable to compare it to previous value so skip it with continue
			continue
		}
		// if anything has been deleted from the original slice, deleted will be > 0
		// for each deletion the slice will be shorter so...
		for d := 1; d <= deleted; d++ { //for each time a value was deleted
			if i == len(bigTerms)-1-d { // if terms[i] would be at the index of (last value) - (number of items deleted)
				finishflag = true // set finishflag to true
				continue          // and skip this value
			}
		}

		if bigTerms[i].Int64() == bigTerms[i-1].Int64() { // compare the value to the one before it and if its the same
			fmt.Printf("deleting term [%d]\n", i)
			bigTerms = append(bigTerms[:i], bigTerms[i+1:]...) // delete it by setting the slice to all values before and after the term to be removed
			fmt.Println(bigTerms)
			//fmt.Println(bigTerms2)

			deleted++ // increment a counter of deleted items
    }
	}
	fmt.Println(bigTerms)
      count := 0
      for _, _ = range(bigTerms){
        count ++
      }
      fmt.Printf("There are %d distinct terms\n", count)
}
