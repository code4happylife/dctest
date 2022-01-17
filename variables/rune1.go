package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str1[2]).Kind())
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str1[2], string(str1[2]))       // 108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è
	fmt.Println("len(str2):", len(str2))        // len(str2)： 8
	testRune := []rune(str2)
	fmt.Println(string(testRune[3]))  
}

/**
uint8
uint8
108 l
232 è
len(str2): 8
言
int32


*/
