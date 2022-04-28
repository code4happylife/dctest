package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
	w := 3.3
	fmt.Printf("w is of type %T\n", w)
	x := 0.123 + 4.334i
	fmt.Printf("x is of type %T\n", x)
	/**
	v is of type int
	w is of type float64
	x is of type complex128
	*/
}
