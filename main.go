package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/symfony-cli/console"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := &console.Application{
		Name:      "DDoser CLI",
		Usage:     "Denial of Service Attacks using Golang",
		Copyright: fmt.Sprintf("(c) 2021-%d GingTeam", time.Now().Year()),
		Version:   "2.0",
		Commands: []*console.Command{
			runCommand,
		},
	}

	app.Run(os.Args)
}
