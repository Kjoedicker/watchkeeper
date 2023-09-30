package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var ports = [4]string{
	"3000",
	"80",
	"8081",
	"22",
}

func checkTCPPort(
	wg *sync.WaitGroup,
	host string,
	port string,
	openPorts *[]string,
	closedPorts *[]string,
) {
	defer wg.Done()

	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		*closedPorts = append(*closedPorts, port)
	} else {
		conn.Close()
		*openPorts = append(*openPorts, port)
	}
}

func displayPorts(status string, ports []string) {
	fmt.Println("  " + status + ":")
	for _, port := range ports {
		fmt.Println("    -" + port)
	}
}

var (
	DEFAULT_STRING = ""
	DEFAULT_INT    = 0
)

func parseCommandLineArguments() (host string, interval int) {
	hostPtr := flag.String("host", DEFAULT_STRING, "Hostname Example: github.com | 140.82.113.4")
	intervalPtr := flag.Int("interval", DEFAULT_INT, "Interval to check connection")

	flag.Parse()

	if *hostPtr == DEFAULT_STRING {
		fmt.Println("A host must be provided")
		os.Exit(1)
	}

	return *hostPtr, *intervalPtr
}

func main() {
	host, interval := parseCommandLineArguments()

WatchLoop:
	for {
		var openPorts []string
		var closedPorts []string

		var wg sync.WaitGroup
		for _, port := range ports {
			wg.Add(1)
			go checkTCPPort(&wg, host, port, &openPorts, &closedPorts)
		}
		wg.Wait()

		fmt.Println(host + " Ports:")
		displayPorts("open", openPorts)
		displayPorts("closed", closedPorts)

		if interval == 0 {
			break WatchLoop
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
