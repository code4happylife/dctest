// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	myfunc(1, 3, 5, 7, 9)
}
