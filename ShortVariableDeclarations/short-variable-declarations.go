package main

import "fmt"

//haha := "hello"
//syntax error: non-declaration statement outside function body

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	//1 2 3 true false no!
}
