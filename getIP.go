package main

import (
	"errors"
	"fmt"
	"net"
)

// This file is used to get machine ip address
func main() {
	ip, err := getClientIp()
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	fmt.Printf("Current ip of test machine: %s\n", ip)
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the ip address is loop back address or not
		if ipaddress, ok := address.(*net.IPNet); ok && !ipaddress.IP.IsLoopback() {
			if ipaddress.IP.To4() != nil {
				return ipaddress.IP.String(), nil
			}
		}
	}
	return "", errors.New("could not find the client ip address")
}
