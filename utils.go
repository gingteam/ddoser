package main

import (
	"bufio"
	"math/rand"
	"os"
)

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
