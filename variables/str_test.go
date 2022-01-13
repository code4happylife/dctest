package main

import "fmt"
//[]byte 切片类型和字符串类型数据相互转换
func main() {
	var data [10]byte
	data[0] = 'T'
	data[1] = 'E'
	var str = string(data[:])
	fmt.Println(str)
	testStr := "hello"
	var a  = []byte(testStr)
	fmt.Printf("%v : %T", a, a)

}

/**

TE
[104 101 108 108 111] : []uint8

*/
