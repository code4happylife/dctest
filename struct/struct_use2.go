package main

import "fmt"

func main() {
	var user struct {
		name string
		age  int
	}
	user.name = "root"
	user.age = 20
	fmt.Printf("%#v\n", user)
	fmt.Printf("%v\n", user)
	/*
		// GoStringer is implemented by any value that has a GoString method,
		// which defines the Go syntax for that value.
		// The GoString method is used to print values passed as an operand
		// to a %#v format.

	*/

}

/**
struct { name string; age int }{name:"root", age:20}
{root 20}


*/
