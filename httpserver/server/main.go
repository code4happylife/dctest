package main

import (
	"log"
	"net/http"
)

const Addr = ":9933"

func main() {
	// 设置http mux路由器
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello)
	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}
	log.Println("Starting server now ...")
	server.ListenAndServe()
}

// sayHello 回调函数
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello, this is a httpserver written in go!</h1>"))
}
