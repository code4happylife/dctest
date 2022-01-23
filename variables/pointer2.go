package main

import "fmt"

//指针

/**
func main() {
	a := 10 // int 类型
	b := &a //内存地址, *int

	//fmt.Println(a, &a)
	//fmt.Println(b)

	fmt.Printf("%v,%p\n", a, &a)
	fmt.Printf("%v,%p\n", b, &b)

	c := *b //根据内存地址去取值
	fmt.Println(c)
}
**/
/**
10,0xc000128058
0xc000128058,0xc00014c018

*/

func modify1(x int) {
	x = 100
}
//指针传值
func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1
	modify2(&a)
	fmt.Println(a)

}

/**

1
100


**/
