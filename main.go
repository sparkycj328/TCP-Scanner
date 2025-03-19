package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

// worker is a function to process the work it receives
func worker(ports chan int, results chan int) {
	var (
		address = "scanme.nmap.org:%s"
	)
	for p := range ports {
		//convert the port number into a string
		intToString := strconv.Itoa(p)

		// dial the port using tcp
		conn, err := net.Dial("tcp", fmt.Sprintf(address, intToString))
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	var (
		successMsg = "%d open\n"
	)
	ports := make(chan int, 1)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// send the port numbers to be scanned on the ports channel
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)

	// sort the ports which were sent back as being open
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf(successMsg, port)
	}
}
