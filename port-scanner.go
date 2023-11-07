package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var _startPort int
	var _endPort int
	var _url, _timeoutStr string

	fmt.Print("Enter Url: ")
	fmt.Scan(&_url)

	fmt.Print("Enter Start Port: ")
	fmt.Scan(&_startPort)

	fmt.Print("Enter End Port: ")
	fmt.Scan(&_endPort)

	fmt.Print("Enter Timeout (Seconds): ")
	fmt.Scan(&_timeoutStr)

	target := _url
	_timeout, _error := time.ParseDuration(_timeoutStr)

	if _error != nil {
		fmt.Print("Invalid Timeout Value.")
		return
	}

	var wg sync.WaitGroup

	for port := _startPort; port <= _endPort; _startPort++ {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s;%d", target, p)
			connection, _error := net.DialTimeout("tcp", address, _timeout)

			if _error != nil {
				fmt.Printf("Port %d is closed\n", p)
			} else {
				defer connection.Close()
				fmt.Printf("Port %d is open\n", p)
			}
		}(port)
	}

	wg.Wait()
}
