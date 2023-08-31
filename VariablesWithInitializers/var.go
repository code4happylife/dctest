package main

import (
	"fmt"
	"reflect"
)

func main(){
	str1:="Golang"
	str2:="Go语言"
	// 字符串是以 byte 数组形式保存的，类型是 uint8，占1个 byte
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	// 打印时需要用 string 进行类型转换，否则打印的是编码值
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n",str2[2],str2[2])
	fmt.Println(reflect.TypeOf(str1[2]).Kind())
	// Go 占 2 byte，语言占 6 byte
	fmt.Println("len(str2): ",len(str2))
}
