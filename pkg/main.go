package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"sync"
	"time"
)

func ScanTarget(IP string, Ports int) {
	var (
		wg        sync.WaitGroup
		sem       = make(chan struct{}, 1000)
		mu        sync.Mutex
		openPorts []int
	)

	for port := 1; port <= Ports; port++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }()

			addr := fmt.Sprintf("%s:%d", IP, p)
			conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			mu.Lock()
			openPorts = append(openPorts, p)
			mu.Unlock()
		}(port)
	}
	wg.Wait()

	sort.Ints(openPorts)
	for i := 0; i < len(openPorts); i++ {
		fmt.Printf("%d	%s\n", openPorts[i], "open")
	}
}

func main() {
	var (
		IPaddr    string
		portRange int
	)

	flag.StringVar(&IPaddr, "ip", "127.0.0.1", "Target IP address")
	flag.IntVar(&portRange, "ports", 1000, "Target Port Range")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if IPaddr == "" {
		fmt.Fprintln(os.Stderr, "Error: --ip is required.\n")
		flag.Usage()
		os.Exit(1)
	}

	log.Println("GScan Scanning Target:", IPaddr)
	fmt.Println("PORT	STATE	SERVICE")
	ScanTarget(IPaddr, portRange)
}
