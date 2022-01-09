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

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func readLineFromFile(fileName string) ([]string, error) {
	var lines []string
	openFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
