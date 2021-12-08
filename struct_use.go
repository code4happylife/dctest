// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Server struct {
	ip        string
	port      string
	username  string
	server_id int
}

func main() {
	var server1 Server
	server1 = Server{ip: "10.10.10.10", port: "1234", username: "test", server_id: 1}
	fmt.Printf("server1 ip: %s\n", server1.ip)
	fmt.Printf("server1 port: %s\n", server1.port)
	fmt.Printf("server1 username: %s\n", server1.username)
	fmt.Printf("server1 server_id: %d\n", server1.server_id)

}
