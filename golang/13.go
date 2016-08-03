package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("13.txt")
	if err != nil {
		panic(err)
	}

	stringData := string(data)
	slicesOfRows := strings.Split(stringData, "\r\n")

	var floatsum float64 = 0

	for i := 0; i < 100; i++ {

		myfloat, err := strconv.ParseFloat(slicesOfRows[i], 64)
		if err != nil {
			panic(err)
		}

		floatsum = floatsum + myfloat
	}

	fmt.Printf(" total of floats is %f\n", floatsum)
}
