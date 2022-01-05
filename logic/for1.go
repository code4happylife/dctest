package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += i
	}
	fmt.Println(sum)

	ret := 0
	for j := 1; j <= 100; j++ {
		ret += j
	}
	fmt.Println("add from 1 to 100,result is:", ret)
}
