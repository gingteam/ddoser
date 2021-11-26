package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

type downloadT struct {
	cli.Helper
	Sock int `cli:"sock" dft:"4"`
}

var downloadCommand = &cli.Command{
	Name: "download",
	Desc: "Download proxy from https://github.com/TheSpeedX/PROXY-List",
	Argv: func() interface{} {
		return new(downloadT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*downloadT)
		runCommand(fmt.Sprintf(
			"wget https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/socks%d.txt",
			argv.Sock,
		))
		ctx.String("Successful download\n")

		return nil
	},
}
