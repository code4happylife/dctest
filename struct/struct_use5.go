package main

import (
	"encoding/json"
	"fmt"
)

// 结构体Tag用法
func main() {
	type student struct {
		Id     string `json:"Id"`
		Gender string `json:"Gender"`
		Email  string `json:"Email"`
		Score  int    `json:"Score"`
	}

	s := student{
		Id:     "001",
		Gender: "M",
		Email:  "student001@gmail.com",
		Score:  600,
	}
	sJson, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("Json Marshal failed:%v", err)
		return
	}
	fmt.Printf("%s", sJson)
}

/*
{"Id":"001","Gender":"M","Email":"student001@gmail.com","Score":600}

*/
