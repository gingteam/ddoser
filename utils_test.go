package main

import (
	"testing"
)

func TestRandom(t *testing.T) {
	seeds := []string{"a", "b", "c"}

	t.Log(random(seeds))
}
