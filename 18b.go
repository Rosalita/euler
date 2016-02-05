package main

import (
  "fmt"
)

// create a struct that defines each node in the triangle
type TriangleNode struct {
  v float64 // the value in the triangle at this node
  l *TriangleNode  // a pointer to another triangleNode, this will be the name of the node on the row below to the left, will be nil for nodes in last row
  r *TriangleNode // a pointer to another triangleNode, this will be the name the node on the row below to the right, will be nil for nodes in last row
  }

func main(){

// input all the triangle values
row0:=  []float64{75}
row1:=  []float64{95, 64}
row2:=  []float64{17, 47, 82}
row3:=  []float64{18, 35, 87, 10}
row4:=  []float64{20, 04, 82, 47, 65}
row5:=  []float64{19, 01, 23, 75, 03, 34}
row6:=  []float64{88, 02, 77, 73, 07, 63, 67}
row7:=  []float64{99, 65, 04, 28, 06, 16, 70, 92}
row8:=  []float64{41, 41, 26, 56, 83, 40, 80, 70, 33}
row9:=  []float64{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
row10:= []float64{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
row11:= []float64{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
row12:= []float64{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
row13:= []float64{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
row14:= []float64{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

//triangle is a slice of slices containing floats
triangle := [][]float64{row0, row1, row2, row3, row4, row5, row6, row7, row8, row9, row10, row11, row12, row13, row14}

// make two new slices to store and the triangle nodes in while they are processed
var nodes []*TriangleNode
var doneNodes []*TriangleNode


// initialise the starting set of triangle nodes for the bottom row of the triangle
for _, value := range triangle[len(triangle)-1]{ // for the last row in the triangle
  newNode := new(TriangleNode) // make a new triangle node
  newNode.v = value   // set the v of this triangle node to the value in the rowslice
  nodes = append(nodes, newNode)  // add the new triangle node to slice of nodes
}


for row := len(triangle) -2; row >= 0; row -- { // length of triangle is 15, row 14 has no linked nodes, so we want to start adding on row 13
  for i, value := range triangle[row]{ // for each value on the row

    newNode := new(TriangleNode) // make a new node
    newNode.v, newNode.l, newNode.r  = value, nodes[i], nodes[i+1] // set its value, then set pointers to the left and right nodes on row below

    if newNode.l.v >= newNode.r.v { // if the value of left node is greater than or equal to right node
        newNode.v += newNode.l.v // add its value to this node
      } else {
        newNode.v += newNode.r.v // otherwise add the other value to this node
      }
    if row == len(triangle) -2{ // if first pass through this loop on the first row
    doneNodes = append(doneNodes, newNode) // store the processed node, slice is empty so append value
    nodes[i] = doneNodes[i] // set next nodes to be processed to the once that have been processed
    } else {
    doneNodes[i] = newNode // store the processed node, by writing it over the top of the previously processed node.
    nodes[i] = doneNodes[i] // set next node to be processed to the once that have just been processed
    }
  }
}

// the answer is the value stored at index 0 of doneNodes
fmt.Println(doneNodes[0].v)
}
