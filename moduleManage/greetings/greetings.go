package greetings

import "fmt"
// Hello returns a greeting for the named person
func Hello(name string) string{
	// Return a greeting that embeds the name in a message
	message:=fmt.Sprintf("Hello,%v.Welcome!",name)
	return message
}
