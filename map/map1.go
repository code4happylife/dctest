package main

import "fmt"

func main() {
	m := map[string]int{"bj": 1, "sh": 2, "gz": 3, "sz": 4, "other": 5}
	for i, v := range m {
		fmt.Println(i, v)
	}
}

/**
sz 4
other 5
bj 1
sh 2
gz 3


 */
