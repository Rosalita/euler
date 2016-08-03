package main

import (
	"bytes"
	"fmt"
	"github.com/cznic/mathutil"
	"strconv"
)

type perm []int

// adding all the methods below make perm satisfy criteria for a sort.interface
func (x perm) Len() int           { return len(x) }
func (x perm) Less(i, j int) bool { return x[i] < x[j] }
func (x perm) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	// find number of permuations for 123456789.
	// there will be !9 permutations which is 362880

	myperm := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // initialise myperm with the lowest value permutation

	var permslice []string // initialise a slice of strings

	var buffer bytes.Buffer // make a new buffer
  sstart := "123456789"

	permslice = append(permslice, sstart)
	for i := 1; i < 10; i++ {  // set to < 10 while testing this will eventually be < 362880
		mathutil.PermutationNext(perm(myperm)) // generate the next permuation

		for _, v := range myperm {
			temp := strconv.Itoa(v) // convert each item in the []int to string
			buffer.WriteString(temp) // store in a buffer
		}
		mystring := buffer.String() // pull strings out of the buffer
		permslice = append(permslice, mystring) // store them in the []string
		buffer.Reset()
	}
	fmt.Println(permslice) // prints all permutations in string format

	// store perms as strings in a slice - done see above

  // for each permuation
	// work through this slice taking bits at the start, multiplying them and seeing if result contains any digits present in the multipliers



}
