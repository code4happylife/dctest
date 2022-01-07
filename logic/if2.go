package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func greater(x, y float64) float64 {
	if tmp := x; tmp > y {
		return tmp
	}
	return y

}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(4, 1, 100),
	)
	fmt.Println(
		greater(1, 100),
		greater(2, 99),
		greater(100, 4),
	)
}
