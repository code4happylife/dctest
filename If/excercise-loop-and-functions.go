package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	//z := x / 2
	z := 1.0
	for itertimes := 1; itertimes <= 10000; itertimes++ {
		z -= (z*z - x) / 2 * z
		if math.Abs(z*z-x) < 0.0001 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
