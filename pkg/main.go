package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		IPaddr    string
		portRange int = 0
	)

	flag.StringVar(&IPaddr, "ip", "127.0.0.1", "Target IP address")
	flag.IntVar(&portRange, "p", 1000, "Target Port Range, Default: Most common 1000 Ports")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if IPaddr == "" {
		fmt.Fprintln(os.Stderr, "Error: --ip is required")
		flag.Usage()
		os.Exit(1)
	}

	log.Println("GScan Scanning Target:", IPaddr)
	fmt.Println("PORT	STATE	SERVICE")
	ScanTarget(IPaddr, portRange)
}
