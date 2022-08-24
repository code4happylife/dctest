package struct_demo

import "fmt"

// Dog 定义Dog结构体
type Dog struct {
	ID   int
	Name string
	Age  int
}

func TestForStruct() {
	// 方式1
	//var dog Dog
	//dog.ID = 0
	//dog.Name = "woki"
	//dog.Age = 3
	// 方式2
	//dog := Dog{ID: 1, Name: "YOYO", Age: 2}
	//方式3
	dog := new(Dog)
	dog.ID = 3
	dog.Name = "YIYI"
	dog.Age = 3

	fmt.Println("dog: ", dog)
}
