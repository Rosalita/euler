package main

import (
  "fmt"
  "github.com/cznic/mathutil"
)

type data []int

// data implements the sort.Interface by defining three methods Len(), Less() and Swap() on the type
// sort.Interface
// type Interface interface {
// Len() int             // Len is the number of elements in a collection
// Less(i, j int) bool   // Less reports true if i is less than j
// Swap(i, j int)        // Swaps the elements with indexes i and j
// }

func (x data) Len() int {return len(x)}
func (x data) Less(i, j int) bool { return x[i] < x[j] }
func (x data) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func main(){
 permdata:= []int{0,1,2,3,4,5,6,7,8,9}
 for i:= 1; i < 1000000; i++{
   mathutil.PermutationNext(data(permdata))
 }
  fmt.Println(permdata)
}
