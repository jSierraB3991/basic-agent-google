package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jSierraB3991/http-cli/portscanner"
)

func main() {
	hostFlag := flag.String("host", "127.0.0.1", "Host to scan (e.g. 127.0.0.1 or localhost)")
	startPortFlag := flag.Int("start", 1, "Start port")
	endPortFlag := flag.Int("end", 1024, "End port")
	timeoutFlag := flag.Duration("timeout", 500*time.Millisecond, "Connection timeout per port")
	concurrencyFlag := flag.Int("concurrency", 500, "Number of concurrent workers")
	commonOnlyFlag := flag.Bool("common", false, "Scan only common well-known ports")

	flag.Parse()

	fmt.Printf("Starting port scanner on host: %s\n", *hostFlag)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful Ctrl+C interruption
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\n[!] Scan interrupted by user. Exiting...")
		cancel()
		os.Exit(0)
	}()

	startTime := time.Now()
	var results []portscanner.Result

	if *commonOnlyFlag {
		fmt.Printf("Scanning %d common ports...\n", len(portscanner.CommonPorts))
		// Extract common ports into a slice or scan them
		var portsToScan []int
		for p := range portscanner.CommonPorts {
			portsToScan = append(portsToScan, p)
		}
		
		// Run customized or concurrent scan for common ports list
		results = make([]portscanner.Result, len(portsToScan))
		for i, p := range portsToScan {
			results[i] = portscanner.ScanPort("tcp", *hostFlag, p, *timeoutFlag)
		}
	} else {
		if *startPortFlag > *endPortFlag {
			fmt.Fprintln(os.Stderr, "Error: Start port cannot be greater than end port.")
			os.Exit(1)
		}
		fmt.Printf("Scanning port range %d to %d with concurrency %d...\n\n", *startPortFlag, *endPortFlag, *concurrencyFlag)
		results = portscanner.ScanRange(ctx, *hostFlag, *startPortFlag, *endPortFlag, *timeoutFlag, *concurrencyFlag)
	}

	duration := time.Since(startTime)

	openCount := 0
	fmt.Println("PORT     STATE    SERVICE")
	fmt.Println("-------------------------")
	for _, res := range results {
		if res.Open {
			openCount++
			fmt.Printf("%-8d OPEN     %s\n", res.Port, res.Service)
		}
	}
	fmt.Println("-------------------------")
	fmt.Printf("Scan completed in %v. Found %d open ports.\n", duration.Round(time.Millisecond), openCount)
}
