package main

import "fmt"

func main()  {
	names:=[4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a:=names[0:2]
	b:=names[1:3]
	fmt.Println(a)
	fmt.Println(b)
	b[0]="xxxx"
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(b)

}

/**
[John Paul George Ringo]
[John Paul]
[Paul George]
[John xxxx George Ringo]
[John xxxx]
[xxxx George]


*/