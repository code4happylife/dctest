package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func greater(x, y float64) float64 {
	if tmp := x; tmp > y {
		return tmp
	} else {
		fmt.Printf("%v <= %v\n", tmp, y)
	}
	return y
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		greater(1, 100),
		greater(-3, 9),
	)
}
/**

27 >= 20
9 20
1 <= 100
-3 <= 9
100 9


*/
