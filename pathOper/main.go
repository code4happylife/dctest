package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var filePath, err = filepath.Abs("./test_dir1/test_dir2/test_dir3/test.log")
	if err != nil {
		log.Fatalf("文件路径解析失败")
		return
	} else {
		// 将文件路径进行分割
		dirs := strings.Split(filePath, string(os.PathSeparator))
		fmt.Println(dirs)
		dir := strings.Join(dirs[0:len(dirs)-1], string(os.PathSeparator))
		fmt.Println(dir)
	}

}
