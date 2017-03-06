package commands

import (
	"fmt"
	"github.com/griff/thonix/api"
	"github.com/griff/thonix/assets"
	"github.com/griff/thonix/server"
	rpcserver "github.com/griff/thonix/rpc/server"
	"github.com/urfave/cli"
	"log"
	"os"
	"net/http"
)


func RunDaemon(c *cli.Context) error {
	s, err := server.NewServer(int32(c.Int("boot-steps")))
	if err != nil {
		return err
	}

	err = s.SetStatusFromString(c.String("state"))
	if err != nil {
		return err
	}

	assetDir := c.String("assets")
	var dir http.FileSystem
	info, err := os.Stat(assetDir)
	if (err != nil && os.IsNotExist(err)) || (err == nil && !info.IsDir()) {
		log.Printf("Serving embedded assets\n")
		dir = assets.AssetFS()
	} else {
		log.Printf("Serving assets from %s\n", assetDir)
		dir = http.Dir(assetDir)
	}

	a := api.API {
		Server: s,
		Assets: dir,
		Shell: c.String("shell"),
	}
	r, err := a.Router()
	if err != nil {
		return err
	}

	http.Handle("/", r)	

	rpcServer, err := rpcserver.NewServer(s)
	err = rpcServer.ListenPublic("")
	if err != nil {
		return err
	}
	err = rpcServer.ListenAdmin("")
	if err != nil {
		rpcServer.Close()
		return err
	}

	port := c.Int("port")
	log.Printf("Application listening on port %d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	
	rpcServer.Close()
	return err
}

func Daemon() cli.Command {
	return cli.Command{
		Name:   "daemon",
		Usage:  "Runs daemon",
		Action: RunDaemon,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port, p",
				Value: 3003,
				Usage: "Port to use for web server",
			},
			cli.StringFlag{
				Name:  "assets",
				Usage: "Path to frontend assets",
				Value: "../frontend/target",
			},
			cli.StringFlag{
				Name:  "tlscert",
				Usage: "Path to TLS certificate file",
			},
			cli.StringFlag{
				Name:  "tlskey",
				Usage: "Path to TLS key file",
			},
			cli.StringFlag{
				Name:  "state",
				Usage: "The state this server is in (on of Booting, BootFailed, Installing and Running)",
				Value: "booting",
			},
			cli.StringFlag{
				Name:  "shell, s",
				Usage: "The shell to use when starting a web term",
				Value: "bash",
			},
			cli.StringFlag{
				Name:  "etc, e",
				Usage: "Path to location of configuration files",
				Value: "/etc/thonix/rescue",
			},
			cli.IntFlag{
				Name:  "boot-steps, c",
				Usage: "The total number of boot steps to expect",
				Value: 0,
			},
		},
	}
}
