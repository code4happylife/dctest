package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}
func main() {
	result := add(1, 2)
	fmt.Println(result)
	res, rem := div(100, 4)
	fmt.Println(res, rem)
}
/*

3
25 0

*/
