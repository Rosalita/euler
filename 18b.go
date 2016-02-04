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
// populate the values of the triangle and the links between each node:

// using a naming convention of r = row, c = column to make node names self referencing.
// create all the nodes in the last row (row 14) first as these nodes don't point to any other nodes.
//04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

r14c0, r14c1, r14c2, r14c3 := &TriangleNode{v:4}, &TriangleNode{v:68}, &TriangleNode{v:98}, &TriangleNode{v:27}
r14c4, r14c5, r14c6, r14c7 := &TriangleNode{v:70}, &TriangleNode{v:98}, &TriangleNode{v:73}, &TriangleNode{v:23}
r14c8, r14c9, r14c10, r14c11 := &TriangleNode{v:9}, &TriangleNode{v:93}, &TriangleNode{v:38}, &TriangleNode{v:53}
r14c12, r14c13, r14c14 := &TriangleNode{v:60}, &TriangleNode{v:04}, &TriangleNode{v:23}

// row 14 can then be defined as a slice containing a pointer to each TriangleNode on that row
//row14 := []*TriangleNode {r14c0, r14c1, r14c2, r14c3, r14c4, r14c5, r14c6, r14c7, r14c8, r14c9, r14c10, r14c11, r14c12, r14c13, r14c14}


fmt.Printf("row 14: %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f \n",
r14c0.v, r14c1.v, r14c2.v, r14c3.v, r14c4.v, r14c5.v,r14c6.v, r14c7.v, r14c8.v, r14c9.v, r14c10.v, r14c11.v, r14c12.v, r14c13.v, r14c14.v)

// create the nodes for row 13, each node on this row will have l and r set to point to a node in the row below
//63 66 04 68 89 53 67 30 73 16 69 87 40 31
r13c0, r13c1, r13c2   := &TriangleNode{v:63, l:r14c0,  r:r14c1},  &TriangleNode{v:66, l:r14c1,  r:r14c2},  &TriangleNode{v:04, l:r14c2,  r:r14c3}
r13c3, r13c4, r13c5   := &TriangleNode{v:68, l:r14c3,  r:r14c4},  &TriangleNode{v:89, l:r14c4,  r:r14c5},  &TriangleNode{v:53, l:r14c5,  r:r14c6}
r13c6, r13c7, r13c8   := &TriangleNode{v:67, l:r14c6,  r:r14c7},  &TriangleNode{v:30, l:r14c7,  r:r14c8},  &TriangleNode{v:73, l:r14c8,  r:r14c9}
r13c9, r13c10, r13c11 := &TriangleNode{v:16, l:r14c9,  r:r14c10}, &TriangleNode{v:69, l:r14c10, r:r14c11}, &TriangleNode{v:87, l:r14c11, r:r14c12}
r13c12, r13c13        := &TriangleNode{v:40, l:r14c12, r:r14c13}, &TriangleNode{v:31, l:r14c13, r:r14c14}

// row 13 can then be defined as a slice containing a pointer to each TriangleNode on that row
row13 := []*TriangleNode {r13c0, r13c1, r13c2, r13c3, r13c4, r13c5, r13c6, r13c7, r13c8, r13c9, r13c10, r13c11, r13c12, r13c13}

fmt.Printf("row 13: %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f \n",
r13c0.v, r13c1.v, r13c2.v, r13c3.v, r13c4.v, r13c5.v,r13c6.v, r13c7.v, r13c8.v, r13c9.v, r13c10.v, r13c11.v, r13c12.v, r13c13.v)

for _, node := range row13 {
  fmt.Println(node.v) // value before adding
  if node.l.v >= node.r.v { // if linked left node value is greater than or equal to linked right node value
    node.v += node.l.v // add the value of the linked left node to this node
    } else {
      node.v += node.r.v // else add the value of the linked right node to this node
    }
  fmt.Println(node.v) // value after adding

}

fmt.Printf("row 13: %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f \n",
r13c0.v, r13c1.v, r13c2.v, r13c3.v, r13c4.v, r13c5.v,r13c6.v, r13c7.v, r13c8.v, r13c9.v, r13c10.v, r13c11.v, r13c12.v, r13c13.v)




}
