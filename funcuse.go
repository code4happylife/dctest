package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	sayhi("you")
}

func sayhi(word string) {
	fmt.Println("nice to meet", word)
}
