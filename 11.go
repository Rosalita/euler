package main

import (
	"fmt"
)

func main() {

	row0 := []int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8}
	row1 := []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0}
	row2 := []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65}
	row3 := []int{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91}
	row4 := []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80}
	row5 := []int{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50}
	row6 := []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70}
	row7 := []int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}
	row8 := []int{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72}
	row9 := []int{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95}
	row10 := []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92}
	row11 := []int{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57}
	row12 := []int{86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58}
	row13 := []int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40}
	row14 := []int{4, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66}
	row15 := []int{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69}
	row16 := []int{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36}
	row17 := []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16}
	row18 := []int{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54}
	row19 := []int{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48}

	grid := [][]int{row0, row1, row2, row3, row4, row5, row6, row7, row8, row9, row10, row11, row12, row13, row14, row15, row16, row17, row18, row19}

	// fmt.Println(grid[0][4]) value at [x][y]
	lrgprod := checkHor(grid)

	b := checkVer(grid)
	if b > lrgprod {
		lrgprod = b
	}
	c := checkDiaSwne(grid)
	if c > lrgprod {
		lrgprod = c
	}

	d := checkDiaNwse(grid)
	if d > lrgprod {
		lrgprod = d
	}

	fmt.Printf("largest product is %d", lrgprod)

}

func checkHor(grid [][]int) int {
	lrgprod := 0
	// for ever row grid[i]
	for row := 0; row < 20; row++ { // for rows 0 - 19
		for col := 0; col < 17; col++ { // for columns 0 - 16
			a, b, c, d := grid[row][col], grid[row][col+1], grid[row][col+2], grid[row][col+3]
			prod := a * b * c * d
			if lrgprod == 0 {
				lrgprod = prod
			} else if prod > lrgprod {
				lrgprod = prod
			}
		}
	}
	return lrgprod
}

func checkVer(grid [][]int) int {
	lrgprod := 0
	for col := 0; col < 20; col++ { // for columns 0 - 19
		for row := 0; row < 16; row++ { // for rows 0 - 16
			a, b, c, d := grid[row][col], grid[row+1][col], grid[row+2][col], grid[row+3][col]
			prod := a * b * c * d
			if lrgprod == 0 {
				lrgprod = prod
			} else if prod > lrgprod {
				lrgprod = prod
			}
		}
	}
	return lrgprod
}

func checkDiaSwne(grid [][]int) int {
	lrgprod := 0
	for row := 3; row < 19; row++ { // for rows 3 - 19
		for col := 0; col < 17; col++ { // for columns 0 - 16
			a, b, c, d := grid[row][col], grid[row-1][col+1], grid[row-2][col+2], grid[row-3][col+3]
			prod := a * b * c * d
			if lrgprod == 0 {
				lrgprod = prod
			} else if prod > lrgprod {
				lrgprod = prod
			}
		}
	}
	return lrgprod
}

func checkDiaNwse(grid [][]int) int {
	lrgprod := 0
	for row := 0; row < 17; row++ { // for rows 0 - 16
		for col := 0; col < 17; col++ { // for columns 0 - 16
			a, b, c, d := grid[row][col], grid[row+1][col+1], grid[row+2][col+2], grid[row+3][col+3]
			prod := a * b * c * d
			if lrgprod == 0 {
				lrgprod = prod
			} else if prod > lrgprod {
				lrgprod = prod
			}
		}
	}
	return lrgprod
}
