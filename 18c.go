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
// create all the nodes in the triangle - using a naming convention of r = row, c = column to make node names self referencing.
//nrow0:= []float64{75}
//nrow1:= []float64{95, 64}
//nrow2:= []float64{17, 47, 82}
//nrow3:= []float64{18, 35, 87, 10}
//nrow4:= []float64{20, 04, 82, 47, 65}
//nrow5:= []float64{19, 01, 23, 75, 03, 34}
//nrow6:= []float64{88, 02, 77, 73, 07, 63, 67}
//nrow7:= []float64{99, 65, 04, 28, 06, 16, 70, 92}
//nrow8:= []float64{41, 41, 26, 56, 83, 40, 80, 70, 33}
//nrow9:= []float64{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
//nrow10:= []float64{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
//nrow11:= []float64{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
//nrow12:= []float64{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
nrow13:= []float64{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
nrow14:= []float64{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

// make a new slice to store the triangle nodes in
var row14nodes []*TriangleNode
var row13nodes []*TriangleNode

for _, value := range nrow14 {
// make a new triangle node
  newNode := new(TriangleNode)
    // set the v of this triangle node to the value in the rowslice
  newNode.v = value
    // store the new triangle node in the new node slice
  row14nodes = append(row14nodes, newNode)
}

fmt.Println(row14nodes)

for i, value := range nrow13 {
  // make a new triangle node
   newNode := new(TriangleNode)
   newNode.v = value
   newNode.l = row14nodes[i]
   newNode.r = row14nodes[i+1]
   if newNode.l.v >= newNode.r.v {
       newNode.v += newNode.l.v
     } else {
       newNode.v += newNode.r.v
     }
//   newNode.l.v =
//   newNode.r.v =
   row13nodes = append(row13nodes, newNode)
}
fmt.Println(row14nodes)
fmt.Println(row13nodes[1].v)


// create all the nodes in the last row (row 14) first because these nodes don't link to any other nodes.
r14c0, r14c1, r14c2, r14c3 := &TriangleNode{v:4}, &TriangleNode{v:62}, &TriangleNode{v:98}, &TriangleNode{v:27}
r14c4, r14c5, r14c6, r14c7 := &TriangleNode{v:23}, &TriangleNode{v:9}, &TriangleNode{v:70}, &TriangleNode{v:98}
r14c8, r14c9, r14c10, r14c11 := &TriangleNode{v:73}, &TriangleNode{v:93}, &TriangleNode{v:38}, &TriangleNode{v:53}
r14c12, r14c13, r14c14 := &TriangleNode{v:60}, &TriangleNode{v:04}, &TriangleNode{v:23}
// create the nodes for row 13, each node on this row will have l and r set to point to a node row 14
r13c0, r13c1, r13c2   := &TriangleNode{v:63, l:r14c0,  r:r14c1},  &TriangleNode{v:66, l:r14c1,  r:r14c2},  &TriangleNode{v:04, l:r14c2,  r:r14c3}
r13c3, r13c4, r13c5   := &TriangleNode{v:68, l:r14c3,  r:r14c4},  &TriangleNode{v:89, l:r14c4,  r:r14c5},  &TriangleNode{v:53, l:r14c5,  r:r14c6}
r13c6, r13c7, r13c8   := &TriangleNode{v:67, l:r14c6,  r:r14c7},  &TriangleNode{v:30, l:r14c7,  r:r14c8},  &TriangleNode{v:73, l:r14c8,  r:r14c9}
r13c9, r13c10, r13c11 := &TriangleNode{v:16, l:r14c9,  r:r14c10}, &TriangleNode{v:69, l:r14c10, r:r14c11}, &TriangleNode{v:87, l:r14c11, r:r14c12}
r13c12, r13c13        := &TriangleNode{v:40, l:r14c12, r:r14c13}, &TriangleNode{v:31, l:r14c13, r:r14c14}
// row 12 nodes
r12c0, r12c1, r12c2   := &TriangleNode{v:91, l:r13c0,  r:r13c1},  &TriangleNode{v:71, l:r13c1,  r:r13c2},  &TriangleNode{v:52, l:r13c2,  r:r13c3}
r12c3, r12c4, r12c5   := &TriangleNode{v:38, l:r13c3,  r:r13c4},  &TriangleNode{v:17, l:r13c4,  r:r13c5},  &TriangleNode{v:14, l:r13c5,  r:r13c6}
r12c6, r12c7, r12c8   := &TriangleNode{v:91, l:r13c6,  r:r13c7},  &TriangleNode{v:43, l:r13c7,  r:r13c8},  &TriangleNode{v:58, l:r13c8,  r:r13c9}
r12c9, r12c10, r12c11 := &TriangleNode{v:50, l:r13c9,  r:r13c10}, &TriangleNode{v:27, l:r13c10, r:r13c11}, &TriangleNode{v:29, l:r13c11, r:r13c12}
r12c12                := &TriangleNode{v:48, l:r13c12, r:r13c13}
//row 11 nodes
r11c0, r11c1, r11c2   := &TriangleNode{v:70, l:r12c0,  r:r12c1},  &TriangleNode{v:11, l:r12c1,  r:r12c2},  &TriangleNode{v:33, l:r12c2,  r:r12c3}
r11c3, r11c4, r11c5   := &TriangleNode{v:28, l:r12c3,  r:r12c4},  &TriangleNode{v:77, l:r12c4,  r:r12c5},  &TriangleNode{v:73, l:r12c5,  r:r12c6}
r11c6, r11c7, r11c8   := &TriangleNode{v:17, l:r12c6,  r:r12c7},  &TriangleNode{v:78, l:r12c7,  r:r12c8},  &TriangleNode{v:39, l:r12c8,  r:r12c9}
r11c9, r11c10, r11c11 := &TriangleNode{v:68, l:r12c9,  r:r12c10}, &TriangleNode{v:17, l:r12c10, r:r12c11}, &TriangleNode{v:57, l:r12c11, r:r12c12}
//row 10 nodes
r10c0, r10c1, r10c2   := &TriangleNode{v:53, l:r11c0,  r:r11c1},  &TriangleNode{v:71, l:r11c1,  r:r11c2},  &TriangleNode{v:44, l:r11c2,  r:r11c3}
r10c3, r10c4, r10c5   := &TriangleNode{v:65, l:r11c3,  r:r11c4},  &TriangleNode{v:25, l:r11c4,  r:r11c5},  &TriangleNode{v:43, l:r11c5,  r:r11c6}
r10c6, r10c7, r10c8   := &TriangleNode{v:91, l:r11c6,  r:r11c7},  &TriangleNode{v:52, l:r11c7,  r:r11c8},  &TriangleNode{v:97, l:r11c8,  r:r11c9}
r10c9, r10c10         := &TriangleNode{v:51, l:r11c9,  r:r11c10}, &TriangleNode{v:14, l:r11c10, r:r11c11}
//row 9 nodes
r9c0, r9c1, r9c2   := &TriangleNode{v:41, l:r10c0, r:r10c1}, &TriangleNode{v:48, l:r10c1, r:r10c2}, &TriangleNode{v:72, l:r10c2, r:r10c3}
r9c3, r9c4, r9c5   := &TriangleNode{v:33, l:r10c3, r:r10c4}, &TriangleNode{v:47, l:r10c4, r:r10c5}, &TriangleNode{v:32, l:r10c5, r:r10c6}
r9c6, r9c7, r9c8   := &TriangleNode{v:37, l:r10c6, r:r10c7}, &TriangleNode{v:16, l:r10c7, r:r10c8}, &TriangleNode{v:94, l:r10c8, r:r10c9}
r9c9               := &TriangleNode{v:29, l:r10c9, r:r10c10}
//row 8 nodes
r8c0, r8c1, r8c2   := &TriangleNode{v:41, l:r9c0,  r:r9c1}, &TriangleNode{v:41, l:r9c1, r:r9c2}, &TriangleNode{v:26, l:r9c2, r:r9c3}
r8c3, r8c4, r8c5   := &TriangleNode{v:56, l:r9c3,  r:r9c4}, &TriangleNode{v:83, l:r9c4, r:r9c5}, &TriangleNode{v:40, l:r9c5, r:r9c6}
r8c6, r8c7, r8c8   := &TriangleNode{v:80, l:r9c6,  r:r9c7}, &TriangleNode{v:70, l:r9c7, r:r9c8}, &TriangleNode{v:33, l:r9c8, r:r9c9}
//row 7 nodes
r7c0, r7c1, r7c2   := &TriangleNode{v:99, l:r8c0,  r:r8c1},  &TriangleNode{v:65, l:r8c1, r:r8c2}, &TriangleNode{v:4, l:r8c2, r:r8c3}
r7c3, r7c4, r7c5   := &TriangleNode{v:28, l:r8c3,  r:r8c4},  &TriangleNode{v:6,  l:r8c4, r:r8c5}, &TriangleNode{v:16, l:r8c5, r:r8c6}
r7c6, r7c7         := &TriangleNode{v:70, l:r8c6,  r:r8c7},  &TriangleNode{v:92, l:r8c7, r:r8c8}
//row 6 nodes
r6c0, r6c1, r6c2   := &TriangleNode{v:88, l:r7c0, r:r7c1}, &TriangleNode{v:2, l:r7c1, r:r7c2}, &TriangleNode{v:77, l:r7c2, r:r7c3}
r6c3, r6c4, r6c5   := &TriangleNode{v:73, l:r7c3, r:r7c4}, &TriangleNode{v:7, l:r7c4, r:r7c5}, &TriangleNode{v:63, l:r7c5, r:r7c6}
r6c6               := &TriangleNode{v:67, l:r7c6, r:r7c7}
//row 5 nodes
r5c0, r5c1, r5c2   := &TriangleNode{v:19, l:r6c0, r:r6c1}, &TriangleNode{v:1, l:r6c1, r:r6c2}, &TriangleNode{v:23, l:r6c2, r:r6c3}
r5c3, r5c4, r5c5   := &TriangleNode{v:75, l:r6c3, r:r6c4}, &TriangleNode{v:3, l:r6c4, r:r6c5}, &TriangleNode{v:34, l:r6c5, r:r6c6}
//row 4 nodes
r4c0, r4c1, r4c2   := &TriangleNode{v:20, l:r5c0, r:r5c1}, &TriangleNode{v:4, l:r5c1, r:r5c2}, &TriangleNode{v:82, l:r5c2, r:r5c3}
r4c3, r4c4         := &TriangleNode{v:47, l:r5c3, r:r5c4}, &TriangleNode{v:65, l:r5c4, r:r5c5}
//row 3 nodes
r3c0, r3c1, r3c2   := &TriangleNode{v:18, l:r4c0, r:r4c1}, &TriangleNode{v:35, l:r4c1, r:r4c2}, &TriangleNode{v:87, l:r4c2, r:r4c3}
r3c3               := &TriangleNode{v:10, l:r4c3, r:r4c4}
//row 2 nodes
r2c0, r2c1, r2c2   := &TriangleNode{v:17, l:r3c0, r:r3c1}, &TriangleNode{v:47, l:r3c1, r:r3c2}, &TriangleNode{v:82, l:r3c2, r:r3c3}
//row 1 nodes
r1c0, r1c1         := &TriangleNode{v:95, l:r2c0, r:r2c1}, &TriangleNode{v:64, l:r2c1, r:r2c2}
//row 0 - only 1 node here at the top of the triangle
r0c0               := &TriangleNode{v:75, l:r1c0, r:r1c1}


// each row can then be defined as a slice containing all the TriangleNodes on that row
row14 := []*TriangleNode {r14c0, r14c1, r14c2, r14c3, r14c4, r14c5, r14c6, r14c7, r14c8, r14c9, r14c10, r14c11, r14c12, r14c13, r14c14}
row13 := []*TriangleNode {r13c0, r13c1, r13c2, r13c3, r13c4, r13c5, r13c6, r13c7, r13c8, r13c9, r13c10, r13c11, r13c12, r13c13}
row12 := []*TriangleNode {r12c0, r12c1, r12c2, r12c3, r12c4, r12c5, r12c6, r12c7, r12c8, r12c9, r12c10, r12c11, r12c12 }
row11 := []*TriangleNode {r11c0, r11c1, r11c2, r11c3, r11c4, r11c5, r11c6, r11c7, r11c8, r11c9, r11c10, r11c11}
row10 := []*TriangleNode {r10c0, r10c1, r10c2, r10c3, r10c4, r10c5, r10c6, r10c7, r10c8, r10c9, r10c10}
row9 := []*TriangleNode {r9c0, r9c1, r9c2, r9c3, r9c4, r9c5, r9c6, r9c7, r9c8, r9c9}
row8 := []*TriangleNode {r8c0, r8c1, r8c2, r8c3, r8c4, r8c5, r8c6, r8c7, r8c8}
row7 := []*TriangleNode {r7c0, r7c1, r7c2, r7c3, r7c4, r7c5, r7c6, r7c7}
row6 := []*TriangleNode {r6c0, r6c1, r6c2, r6c3, r6c4, r6c5, r6c6}
row5 := []*TriangleNode {r5c0, r5c1, r5c2, r5c3, r5c4, r5c5}
row4 := []*TriangleNode {r4c0, r4c1, r4c2, r4c3, r4c4}
row3 := []*TriangleNode {r3c0, r3c1, r3c2, r3c3}
row2 := []*TriangleNode {r2c0, r2c1, r2c2}
row1 := []*TriangleNode {r1c0, r1c1}
row0 := []*TriangleNode {r0c0}

// the whole triangle can then be defined as a slice containing all the rows
triangle := [][]*TriangleNode{row0, row1, row2, row3, row4, row5, row6, row7, row8, row9, row10, row11, row12, row13, row14}

// triangle[0] is the top of the triangle it is a slice containing a pointer to the TriangleNode at the very top
// triangle[1] is the next row of the triangle, it is a slice containing two pointer to both the TriangleNodes on this row
// triangle[2] a slice containing three triangle nodes... etc..



for row := len(triangle) -2; row >= 0; row -- { // length of triangle is 15, row 14 has no linked nodes, so we want to start adding on row 13
  for _, node := range triangle[row]{
    if node.l.v >= node.r.v { // if linked left node value is greater than or equal to linked right node value
      node.v += node.l.v // add the value of the linked left node to this node
      } else {
        node.v += node.r.v // else add the value of the linked right node to this node
      }
  }
}

// after adding all rows, the answer is the value stored in the top node of the triangle
fmt.Println(r0c0.v)
}
