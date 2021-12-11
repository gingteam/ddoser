package main

import (
	"github.com/mkideal/cli"
)

var downloadCommand = &cli.Command{
	Name: "download",
	Desc: "Download proxy from https://github.com/TheSpeedX/PROXY-List",
	Fn: func(ctx *cli.Context) error {
		runCommand("wget https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/socks4.txt")
		ctx.String("Successful download\n")

		return nil
	},
}
