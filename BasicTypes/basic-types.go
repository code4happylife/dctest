package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type:%T Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T Value:%v\n", z, z)
	/*
		Type:bool Value:false
		Type:uint Value:18446744073709551615
		Type:complex128 Value:(2+3i)
	**/
}
