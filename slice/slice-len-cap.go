package main

import "fmt"

func main(){
	s:=[]int{2,3,5,7,11,13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s=s[:4]
	printSlice(s)
	s=s[2:]
	printSlice(s)
	//s=s[:10]
	//printSlice(s)
	//panic: runtime error: slice bounds out of range [:10] with capacity 4
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}
/*
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]

**/