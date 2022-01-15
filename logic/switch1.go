package main

import "fmt"

func main() {
	var a string = "green"
	switch a {
	case "green":
		fmt.Println("The color is green")
		fallthrough
	case "red":
		fmt.Println("The color can be red also")
		fallthrough
	case "yellow":
		fmt.Println("Why not choose yellow")
		fallthrough
	default:
		fmt.Println("You bet it")

	}
}

/**
The color is green
The color can be red also
Why not choose yellow
You bet it


*/
