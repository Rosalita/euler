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
	init := big.NewInt(int64(0))
	var bigTerms bigIntSlice
	uniqueTerms := bigIntSlice{init} // slice of pointers to Int type from math/big package

	for i := start; i <= end; i++ { // loop through all integers
		for j := start; j <= end; j++ { // loop through all powers
      big1, big2, bigAns:= big.NewInt(int64(i)), big.NewInt(int64(j)), big.NewInt(int64(1))
      bigAns.Exp(big1, big2, nil)
      bigTerms = append(bigTerms, bigAns)
		}
	}

	isPresent := false
	for i, v := range uniqueTerms{ // check each value in the slice of bigTerms to see if bigAns is already there
		fmt.Printf("uniqueTerm is %d\n", v)
	   for i2, v2 := range bigTerms{
			 fmt.Printf("bigTerm is %d\n", v2)
			 if uniqueTerms[i] == bigTerms[i2] {
				 isPresent = true
			 }
		   if isPresent != true {
		     uniqueTerms = append(uniqueTerms, bigTerms[i2])

		     isPresent = false
		     continue
		   }
		 }
	}
fmt.Println(bigTerms)
fmt.Println(uniqueTerms)
sort.Sort(bigIntSlice(bigTerms)) // sort the terms into numerical order

//	fmt.Println(bigTerms)
      count := 0
      for _, _ = range(bigTerms){
        count ++
      }
      fmt.Printf("There are %d distinct terms\n", count -1) // don't count the 0 which was used to initialise the slice
}
