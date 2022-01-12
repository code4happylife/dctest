package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func WebSocketClient() {
	url := "ws://localhost:7788/websocket"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("connection fail:", err)
	}
	log.Printf("response is:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, []byte("Hello, I am a visitor to this websocket"))
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Receive message: %s", message)

	}
	<-done
}

func main() {
	WebSocketClient()
}
