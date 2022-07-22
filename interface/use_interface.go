package main

import (
	"fmt"
	"reflect"
)

type Sayer interface {
	Say() string
}

type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

type Horse struct{}

func (h Horse) Say() string { return "neigh" }

type Tiger struct{}

func (t Tiger) Say() string { return "woaw" }

func main() {
	c := Cat{}
	fmt.Println("Cat says:", c.Say())
	d := Dog{}
	fmt.Println("Dog says:", d.Say())
	t := Tiger{}
	fmt.Println("Tiger says:", t.Say())

	animals := []Sayer{c, d, t}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}

	animals = append(animals, Horse{})
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}

	MakeCatTalk(c)
	MakeTalk(animals[2])
	MakeTalk(animals[3])
}

func MakeCatTalk(c Cat) {
	fmt.Println("Cat says:", c.Say())
}

func MakeTalk(s Sayer) {
	fmt.Println(reflect.TypeOf(s).Name(), "says:", s.Say())
}
