package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/mkideal/cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := cli.Root(
		rootCommand,
		cli.Tree(attackCommand),
		cli.Tree(downloadCommand),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
