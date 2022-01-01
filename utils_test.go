package main

import (
	"net/http"
	"testing"
)

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

func TestProxy(t *testing.T) {
	conn, err := connect("141.94.26.29:3129", "http://httpbin.org/get")

	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)

	if err != nil {
		t.Error(err)
	}

	req.Write(conn)
}
