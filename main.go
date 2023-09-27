package main

import (
	"fmt"
	"net"
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

func main() {
	// TODO: pull this from a command line argument
	host := ""

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
}
