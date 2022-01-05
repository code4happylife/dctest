package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//ch := make(chan bool)
	for i := 1; i <= 500; i++ {
		file_name := "./test" + strconv.Itoa(i) + ".txt"
		fmt.Printf("%v", file_name)
		f, err := os.Open(file_name)
		if err != nil {
			fmt.Printf("%v", err)
		}
		b1 := make([]byte, 10)
		n1, err := f.Read(b1)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("%v", n1)
	}
	/**
	go func() {
		for true {
			select {
			case <-ch:
				fmt.Println("get data")
			}
		}
	}()

	go func() {
		for {
			ch <- true
			time.Sleep(3 * time.Second)
		}
	}()

	select {}

	*/
	for {
		time.Sleep(3 * time.Second)
	}
}
