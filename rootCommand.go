package main

import (
	"fmt"

	"github.com/labstack/gommon/color"
	"github.com/mkideal/cli"
)

type appT struct {
	Name    string
	Version string
}

var app = appT{
	"Go DDoser",
	"v1.0.0",
}

func sprintHeader() string {
	return fmt.Sprintf("%s version %s", color.Green(app.Name), color.Yellow(app.Version))
}

func printHeader() {
	fmt.Print(sprintHeader() + "\n\n")
}

type rootT struct {
	cli.Helper
}

var rootCommand = &cli.Command{
	Desc: sprintHeader(),
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		fmt.Print("Usage: go-ddoser <command>\n")
		fmt.Printf("More information on usage with %s flag.\n", color.Yellow("--help"))
		fmt.Printf("Backed by %s\n", color.Green("GingTeam"))

		return nil
	},
}
