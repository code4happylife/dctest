// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
	Phone   string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203", "111"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101", "222"}
	personDB["0"] = PersonInfo{"0", "Kendy", "Room 001", "333"}

	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person 1234", person.Name)
	} else {
		fmt.Println("Did not found person with ID 1234")
	}

	person, ok = personDB["0"]
	if ok {
		fmt.Println("Found person with ID 0")
	} else {
		fmt.Println("Did not found person with ID 0")
	}
}
