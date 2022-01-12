package main

import (
	"errors"
	"fmt"
)

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error:name is null")
	}
	fmt.Println("Hello ", name)
	return nil
}

func sayHi(name string) error {
	if len(name) == 0 {
		return errors.New("should say Hi to not null person")
	}
	fmt.Println("Hi", name)
	return nil
}
func main() {
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
	if err := sayHi("Jude"); err != nil {
		fmt.Println(err)
	}
}

//error:name is null
