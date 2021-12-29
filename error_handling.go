// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func test_result(a, b int) (bool, error) {
	if a > b {
		return true, nil
	}

	return false, errors.New("Not greater than")

}

func main() {
	v1, v2 := test_result(3, 1)
	fmt.Printf("result is %v\n", v1)
	fmt.Printf("error is %v\n", v2)
}
