package main

import "fmt"

type Set map[string]struct{}

func main() {
	set := make(Set)

	for _, item := range []string{"A", "A", "B", "C", "A"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) // 3
	if oo, ok := set["A"]; ok {
		fmt.Println("A exists") // A exists
		fmt.Println(oo) //{}
	}

}
/**

3
A exists
{}


*/
