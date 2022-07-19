package main

import "fmt"

var m float32 = 55
var (
	name    string = "Ross"
	phone   string = "0001122334455"
	address string = "DownTown"
)

func main() {
	var i int
	i = 42
	fmt.Println(i)
	var j int = 27
	fmt.Println(j)
	k := 12
	fmt.Println(k)
	fmt.Printf("%v %T\n", m, m)
	fmt.Println("This guy named", name, "'s phone number is ", phone, "lives in", address)
}
