package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y float64 = 2, 1000
	var total int64 = 0
	power := math.Pow(x, y)
	fmt.Printf("for the power %.f \n", power)
	// convert float to a string
	stringpower := strconv.FormatFloat(power, 'f', 0, 64) // [float, format (e, E, f, g, G), precision, bitsize]

	for _, v := range stringpower {
		numString := string(v)
		numInt64, err := strconv.ParseInt(numString, 10, 64)

		if err != nil {
			panic(err)
		}

		total = total + numInt64
	}
	fmt.Printf("total power digit sum is %d", total)
}
