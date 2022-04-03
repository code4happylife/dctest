package main

import (
	"example.com/greetings"
	"fmt"
)

func main(){
	message :=greetings.Hello("Robot")
	fmt.Println(message)
}
