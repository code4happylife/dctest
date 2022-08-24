package main

import (
	"dctest/struct_demo"
	"fmt"
	"os/exec"
)

func main() {
	command := "ifconfig"
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	//	test struct
	struct_demo.TestForStruct()
}
