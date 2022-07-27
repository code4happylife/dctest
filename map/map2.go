package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[string]string{"one": "No.1"}
	m2 := map[string]string{"one": "No.1"}
	if reflect.DeepEqual(m1, m2) {
		fmt.Println("m1 and m2 are equal.")
	}

}
