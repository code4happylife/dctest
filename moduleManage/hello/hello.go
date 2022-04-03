package main

import (
	"example.com/greetings"
	"fmt"
)

func main(){
	message :=greetings.Hello("Robot")
	fmt.Println(message)
}

/*

go mod edit -replace example.com/greetings=../greetings
go mod tidy
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000

**/
