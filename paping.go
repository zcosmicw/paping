package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(cyan + "made by zcosmicw | Usage: paping <host> <port>" + reset)
		return
	}

	host := os.Args[1]
	port := os.Args[2]
	address := net.JoinHostPort(host, port)

	fmt.Printf(cyan+"Pinging %s on TCP port %s...\n\n"+reset, host, port)

	for {
		start := time.Now()
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf(red+"🔴 Connection failed to %s on TCP port %s (%v)"+reset+"\n", host, port, err)
		} else {
			fmt.Printf(green+"🟢 Connected to %s on TCP port %s in %v"+reset+"\n", host, port, duration)
			conn.Close()
		}

		time.Sleep(1 * time.Second)
	}
}
