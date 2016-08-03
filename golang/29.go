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
	var start, end int = 2, 100

	var bigTerms bigIntSlice
	var uniqueTerms bigIntSlice

	for i := start; i <= end; i++ { // loop through all integers
		for j := start; j <= end; j++ { // loop through all powers
      big1, big2, bigAns:= big.NewInt(int64(i)), big.NewInt(int64(j)), big.NewInt(int64(1))
      bigAns.Exp(big1, big2, nil)
			fmt.Printf(" %d ^ %d = big Ans is %d\n", big1, big2, bigAns)
      bigTerms = append(bigTerms, bigAns)
		}
	}

for _, v := range bigTerms{
	if !contains(uniqueTerms, v){
		uniqueTerms = append(uniqueTerms, v)
	}
}

sort.Sort(bigIntSlice(bigTerms)) // sort the terms into numerical order
//fmt.Println(bigTerms) // prints all terms found
//fmt.Println()
//fmt.Println(uniqueTerms) // prints unique terms
//fmt.Println()
//sort.Sort(bigIntSlice(uniqueTerms))
//fmt.Println(uniqueTerms) // prints unique terms

      count := 0
      for i, v := range(uniqueTerms){
				fmt.Println(i, v)
        count ++
      }
      fmt.Printf("There are %d distinct terms\n", count)

}
func contains(slice []*big.Int, x *big.Int) bool {
    for _, v := range slice {
        if v.Int64() == x.Int64() {
            return true
        }
    }
    return false
}
