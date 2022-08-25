package main

import (
	"dctest/goroutine"
	"time"
)

func main() {
	//command := "ifconfig"
	//cmd := exec.Command("/bin/bash", "-c", command)
	//
	//output, err := cmd.Output()
	//if err != nil {
	//	fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
	//	return
	//}
	//fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	////	test struct
	//struct_demo.TestForStruct()
	//fmt.Printf("cpu num = %d\n", runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//go goroutine.Loop()
	//go goroutine.Loop1()
	//time.Sleep(60 * time.Second)

	go goroutine.Send()
	go goroutine.Receive()
	time.Sleep(60 * time.Second)
}
