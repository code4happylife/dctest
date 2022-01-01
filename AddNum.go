package main

import (
	"errors"
	"fmt"
)

func AddNum(a, b int) (res int, err error) {

	if a < 0 || b < 0 {
		err = errors.New("Only accept positive values")
		return
	}
	return a + b, nil
}

func main() {
	res, err := AddNum(1, 3)
	if err != nil {
		fmt.Printf("error")
	}
	fmt.Printf("%v", res)
}
