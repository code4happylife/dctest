package main

import "fmt"

type baby struct{}

type adult struct{}

// 分别给baby类型和adult类型定义cry()方法
func (b baby) cry() {
	fmt.Println("cry for hunger")
}

func (a adult) cry() {
	fmt.Println("cry for unhappy things")
}

// 给emotion这个接口 类型定义cry()方法
type emotion interface {
	cry()
}

// interface （接口）是一种特殊的类型
// 实现了接口中定义的方法的类型，可以与接口类型一致
func main() {

	b := baby{}
	var e emotion
	e = b
	e.cry()

	a := adult{}
	e = a
	e.cry()
}

/**
cry for hunger
cry for unhappy things


*/
