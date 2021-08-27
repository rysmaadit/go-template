package main

import (
	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
