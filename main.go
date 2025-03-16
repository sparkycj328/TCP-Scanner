package main

import (
	"fmt"
	"net"
)

func main() {
	if _, err := net.Dial("tcp", "scanme.nmap.org:80"); err == nil {
		fmt.Println("Connection successful!")
	} else {
		fmt.Println(err.Error())
	}
}
