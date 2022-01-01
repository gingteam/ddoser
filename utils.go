package main

import (
	"bufio"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

// Returns an array of User-agent
func getUserAgents(number int) []string {
	ua := make([]string, number)

	for i := range ua {
		ua[i] = browser.Random()
	}

	return ua
}

// Returns a random element in the []string
func random(seeds []string) string {
	return seeds[rand.Intn(len(seeds))]
}

func readLines(fileName string) []string {
	var lines []string
	openFile, _ := os.Open(fileName)
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func connect(proxy, address string) (net.Conn, error) {
	netDialer := net.Dialer{
		Timeout: time.Second * 5,
	}

	conn, err := netDialer.Dial("tcp", proxy)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("CONNECT", address, nil)

	if err != nil {
		return nil, err
	}

	// try to write the request
	err = req.Write(conn)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
