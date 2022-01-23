package main

import "fmt"

func main() {
	type person struct {
		name string
		id   int
	}
	var p = new(person)
	p.id = 1
	p.name = "Root"

	fmt.Printf("%#v", p)
}

/**
&main.person{name:"Root", id:1}

*/
