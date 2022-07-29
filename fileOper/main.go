package main

import (
	"fmt"
	"os"
)

func main() {
	workSpace, _ := os.Getwd() // 获取项目的根目录
	fmt.Println(workSpace)

}
