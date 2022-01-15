package main

import (
	"log"
)

func main() {
	temp := map[string]string{"sleep": "nice", "eat": "great"}
	if temp["sleep"] == "nice" {
		log.Printf("great")

	}
}

/**

2022/01/15 11:09:19 great

 */
