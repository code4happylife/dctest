package main

import (
	"fmt"
	"log"
	"net"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	panic("unable to determine locla ip")
}
func main() {
	ip, err := GetLocalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
