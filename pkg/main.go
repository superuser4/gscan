package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var (
		ipAddr    string
		portRange int
		timeoutMs int
		workers   int
	)

	flag.StringVar(&ipAddr, "ip", "127.0.0.1", "Target IP address (required)")
	flag.IntVar(&portRange, "p", 0, "Max port number to scan (Most common 1000)")
	flag.IntVar(&timeoutMs, "t", 1000, "Timeout in milliseconds (default 500)")
	flag.IntVar(&workers, "w", 100, "Number of concurrent workers (default 200)")
	flag.Parse()

	if ipAddr == "" {
		fmt.Fprintln(os.Stderr, "Error: --ip is required")
		flag.Usage()
		os.Exit(1)
	}

	log.Printf("GScan scanning %s ports 1-%d with %d workers, %dms timeout\n",
		ipAddr, portRange, workers, timeoutMs)
	fmt.Println("PORT	STATE	SERVICE")

	scanner := &PortScanner{
		IP:           ipAddr,
		MaxPort:      portRange,
		Timeout:      time.Duration(timeoutMs) * time.Millisecond,
		WorkerAmount: workers,
	}
	scanner.ScanTarget()
}
