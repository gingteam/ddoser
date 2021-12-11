package main

import "testing"

func TestRandom(t *testing.T) {
	seeds := []string{"a", "b", "c"}

	t.Log(random(seeds))
}

func TestUserAgent(t *testing.T) {
	ua := getUserAgents(500)
	if len(ua) != 500 {
		t.Error("This error cannot happen")
	}

	t.Log(ua[0])
}
