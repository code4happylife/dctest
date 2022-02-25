package main

import "fmt"

func main(){
	a:=[]int{1,2,3}
	b:=[]int{1,2,3}
	if a==b{
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}
}
/*
.\main.go:8:6: invalid operation: a == b (slice can only be compared to nil)


*/
