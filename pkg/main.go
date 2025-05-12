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
		timeMs int
	)

	scanner := &PortScanner{}

	flag.StringVar(&scanner.IP, "ip", "", "Target IP address (required)")
	flag.IntVar(&scanner.MaxPort, "p", 0, "Max port number to scan (Most common 1000)")
	flag.IntVar(&timeMs, "t", 500, "Timeout in milliseconds")
	flag.IntVar(&scanner.WorkerAmount, "w", 100, "Number of concurrent workers")
	flag.Parse()

	if scanner.IP == "" {
		fmt.Fprintln(os.Stderr, "Error: --ip is required")
		flag.Usage()
		os.Exit(1)
	}
	scanner.Timeout = time.Duration(timeMs) * time.Millisecond
	log.Printf("GScan scanning %s ports with %d workers, %dms timeout\n",
		scanner.IP, scanner.MaxPort, scanner.WorkerAmount, timeMs)
	fmt.Println("PORT	STATE	SERVICE")

	scanner.ScanTarget()
}
