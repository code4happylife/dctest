package goroutine

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 10)
var timeout = make(chan bool)

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 5)
		fmt.Printf("%d,", i)
	}
}

func Loop1() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d,", i)
	}
}

// 发送数据
func Send() {
	time.Sleep(1 * time.Second)
	chanInt <- 1
	time.Sleep(1 * time.Second)
	chanInt <- 2
	time.Sleep(1 * time.Second)
	chanInt <- 3
	time.Sleep(2 * time.Second)
	timeout <- true

}

//接收数据
func Receive() {
	//num := <-chanInt
	//fmt.Println("num:", num)
	//num = <-chanInt
	//fmt.Println("num:", num)
	//num = <-chanInt
	//fmt.Println("num:", num)
	for {
		select {
		case num := <-chanInt:
			fmt.Println("num:", num)
		case <-timeout:
			fmt.Println("timeout...")
		}
	}

}
