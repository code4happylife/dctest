package main

import "fmt"

func main(){
	var s []int
	fmt.Println(s,len(s),cap(s))
	if s==nil{
		fmt.Println("This is nil slice!")
	}
}

/**
[] 0 0
This is nil slice!


**/