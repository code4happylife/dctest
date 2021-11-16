package main

import (
	"flag"
	"fmt"
	"os/exec"
)
//This file is used to execute some commands some times according to what you pass to the command-line flags
//If you go build this filt into an executeble file with the name tools. Then you can use this tool like this 
//for example: ./tools -command_to_test="date" -execute_times="3"
//var testName = flag.String("name","tom","please input your name")
//var testAge = flag.Int("age",20,"input your age")
//var flag_var int
var command_test string
var execute_times int

func Init() {
	//flag.IntVar(&flag_var,"flagname",1234,"help message for flagname")
	flag.StringVar(&command_test,"command_to_test","ls","choose a command to execute in bash")
	flag.IntVar(&execute_times,"execute_times",10,"execute the specific command for some times")

}

func main(){
	Init()
	flag.Parse()
	for i:=0;i<=execute_times;i++{
		cmd := exec.Command(command_test)
		output,err := cmd.Output()
		if err != nil{
			fmt.Printf("Execute command:%s failed with error:%s in iteration:%d",command_test,err.Error(),i)
		}
		fmt.Printf("Execute command:%s successfully with output:%s in iteration:%d",command_test,string(output),i)
	}



	fmt.Println("command_to_test=",command_test)
	fmt.Println("execute_times=",execute_times)
}
