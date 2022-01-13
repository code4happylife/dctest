package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func touchFile(fileContent, fileFullPathName string) bool {
	lock := &sync.RWMutex{}
	lock.Lock()
	file, err := os.Create(fileFullPathName)
	if err != nil {
		return false
	}
	_, err = file.Write([]byte(fileContent))
	if err != nil {
		return false
	}
	lock.Unlock()
	log.Println("File created successfully!")
	return true
}

func main() {
	ret := touchFile("hello world", "./test2/test.txt")
	if ret != true {
		fmt.Printf("Failed to create test file")
		return
	}
	fmt.Printf("File created successfully")
}

/*
2022/01/13 23:03:48 File created successfully!
File created successfully


*/
