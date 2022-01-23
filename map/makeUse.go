package main

import "fmt"

func main() {
	var a *int
	a = new(int)
	*a = 3
	fmt.Println(*a)
	var b map[string]int
	b = make(map[string]int, 10)
	b["ook"] = 1
	fmt.Println(b)
}

/*

make 初始化map
修改前:
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x108a136]

goroutine 1 [running]:
main.main()

panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()


修改后：
3
map[ook:1]


**/
