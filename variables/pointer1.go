package main

import "fmt"

//指针
func main() {
	a := 10 // int 类型
	b := &a //内存地址, *int

	//fmt.Println(a, &a)
	//fmt.Println(b)

	fmt.Printf("%v,%p\n", a, &a)
	fmt.Printf("%v,%p\n", b, &b)
}

/**
10,0xc000128058
0xc000128058,0xc00014c018

*/
