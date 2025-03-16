package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var (
		address    = "scanme.nmap.org:%s"
		successMsg = "%s open\n"
	)
	for i := 1; i <= 1024; i++ {
		s := strconv.Itoa(i)
		conn, err := net.Dial("tcp", fmt.Sprintf(address, s))
		if err != nil {
			// port is closed or filtered.
			continue
		}
		// close the successful connection
		conn.Close()

		fmt.Printf(successMsg, s)
	}
}
