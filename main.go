package main

import (
	"github.com/griff/thonix/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		commands.Daemon(),
		commands.Report(),
		commands.BootFailed(),
	}
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(-1)
	}
}
