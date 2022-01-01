// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len of mySlice:", len(mySlice))
	fmt.Println("cap of mySlice", cap(mySlice))
}
