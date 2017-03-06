package commands

import (
	"errors"
	"fmt"
	"github.com/griff/thonix-backend/dmesg"
	"github.com/urfave/cli"
)

func ReportAction(c *cli.Context) error {
	block := c.Args().Get(0)
	if block == "" {
		return errors.New("Missing block name")
	}
	_, err := dmesg.Write(fmt.Sprintf("thonix: %s ", block))
	return err
}

func Report() cli.Command {
	return cli.Command{
		Name:      "report",
		Usage:     "Reports boot progress",
		Action:    ReportAction,
		ArgsUsage: "[block]",
		Flags:     []cli.Flag{},
	}
}
