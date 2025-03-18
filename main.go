package main

import (
	"fmt"
	"sync"
)

// worker is a function to process the work it receives
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	//var (
	//	address    = "scanme.nmap.org:%s"
	//	successMsg = "%s open\n"
	//)
	var wg sync.WaitGroup

	ports := make(chan int, 1)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
