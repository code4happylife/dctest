package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >=%g\n", v, lim)
	}
	return lim
	//return v //undefined: v
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))

	//9
	//27 >=10
	//10

}
