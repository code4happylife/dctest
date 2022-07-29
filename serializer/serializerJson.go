package main

import (
	"encoding/json"
	"fmt"
)

type Computer struct {
	Display  string `json:"display"`
	Keyboard string `json:"keyboard"`
	Mouse    string `json:"mouse"`
}

func (c Computer) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

func main() {
	c := Computer{
		Display:  "ThinkVision",
		Keyboard: "HP",
		Mouse:    "HP",
	}
	fmt.Printf("%s", c) //Because struct Computer has already implement String method.
}

/**
{"display":"ThinkVision","keyboard":"HP","mouse":"HP"}

*/
