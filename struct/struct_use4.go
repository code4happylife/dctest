package main

import "fmt"

func main() {
	type test struct {
		a int
		b int
		c int
		d int
	}
	var ret = new(test)
	ret.a = 1
	ret.b = 2
	ret.c = 3
	ret.d = 4

	fmt.Printf("%p\n", &ret.a)
	fmt.Printf("%p\n", &ret.b)
	fmt.Printf("%p\n", &ret.c)
	fmt.Printf("%p\n", &ret.d)
}

/**
内存对齐：
0xc0000b4020
0xc0000b4028
0xc0000b4030
0xc0000b4038


*/
