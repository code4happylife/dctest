package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

//open test.txt: no such file or directory
