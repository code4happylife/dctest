package main

import "fmt"

func exchange(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var a, b int
	a = 2
	b = 3
	fmt.Println(a, b)
	exchange(&a, &b)
	fmt.Println(a, b)
	c := 10
	d := 20
	fmt.Println(c, d)
	c, d = d, c
	fmt.Println(c, d)
}
