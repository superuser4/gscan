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

	flag.StringVar(&ipAddr, "ip", "", "Target IP address (required)")
	flag.IntVar(&portRange, "p", 1000, "Max port number to scan (default 1000)")
	flag.IntVar(&timeoutMs, "t", 500, "Timeout in milliseconds (default 500)")
	flag.IntVar(&workers, "w", 200, "Number of concurrent workers (default 200)")
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
