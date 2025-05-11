package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func ScanTarget(IP string, Ports int) {
	var wg sync.WaitGroup

	for port := 1; port <= Ports; port++ {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", IP, p), 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d	open\n", p)
		}(port)
	}
	wg.Wait()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No Target IP provided! Quitting...")
		os.Exit(1)
	}
	var IPaddr string = os.Args[1]
	log.Println("GScan Scanning Target:", IPaddr)
	fmt.Println("PORT	STATE	SERVICE")
	ScanTarget(IPaddr, 30000)
}
