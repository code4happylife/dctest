package main

import (
	"fmt"
)
//append 向slice 添加元素
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/**
[0 0 0 0 0 1 2 3]

*/
