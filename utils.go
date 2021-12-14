package main

import (
	"bufio"
	"math/rand"
	"os"
	"os/exec"

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

func runCommand(command string) {
	exec.Command("sh", "-c", command).Run()
}
