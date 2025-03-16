package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	var (
		address    = "scanme.nmap.org:%s"
		successMsg = "%s open\n"
		wg         sync.WaitGroup
	)
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			// defer the Wait gorup document until the goroutine is finished executing
			defer wg.Done()

			//convert the port number into a string
			s := strconv.Itoa(j)

			// dial the port using tcp
			conn, err := net.Dial("tcp", fmt.Sprintf(address, s))
			if err != nil {
				// port is closed or filtered.
				return
			}
			// close the successful connection
			conn.Close()
			fmt.Printf(successMsg, s)
		}(i)
	}
	wg.Wait()
}
