package main

import (
	"fmt"
)

func main() {
	x, y := 0, 0
	fmt.Println("what is the value of x?")
	fmt.Scanln(&x)
	fmt.Println("what is the value of y?")
	fmt.Scanln(&y)
  fx := float64(x)
	fy := float64(y)

	fmt.Println(solveAnnoyingMaths(fx, fy))

}

func solveAnnoyingMaths(x float64, y float64) float64{

  xfac:= findFactorial(x)
	yfac:= findFactorial(y)
	total := x + y
	totalfac:= findFactorial(total)
  xyfac:= xfac * yfac
	answer:= totalfac/xyfac
		return answer
}

func findFactorial(x float64) float64{
		 var i, fac float64 = 1, 1
			for i = 1; i <= x; i++ {
			fac = fac * i
			}
			return fac
}
