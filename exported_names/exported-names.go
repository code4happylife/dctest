package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.pi)
	//./exported-names.go:7:14: cannot refer to unexported name math.pi
	fmt.Println(math.Pi)
	//3.141592653589793
}
