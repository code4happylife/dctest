// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var stockcode int32 = 123
	var nowdate string = "20211208"
	var result = "code is %d and now date is %s"
	var specific_value = fmt.Sprintf(result, stockcode, nowdate)
	fmt.Println(specific_value)
}
