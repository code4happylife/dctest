package main

import "fmt"

func main() {
	var a = new(int) //a=nil
	//func new(Type) *Type
	*a = 100
	fmt.Println(*a)

	var b = make(map[string]int)
	b["aa"] = 1
	fmt.Println(b)
}

/*
100
map[aa:1]


**/
