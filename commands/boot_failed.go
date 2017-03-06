package commands

import (
  "github.com/griff/thonix-backend/rpc"
  "github.com/urfave/cli"
  "golang.org/x/net/context"
)

func BootFailedAction(c *cli.Context) error {
  rpcClient, err := rpc.ConnectAdmin(context.Background(), "")
  if err != nil {
    return err
  }
  defer rpcClient.Close()

  _, err = rpcClient.SetState(func (params rpc.ServerAdmin_setState_Params) error {
    params.SetNewState(rpc.Status_bootFailed)
    return nil
  }).Struct()
  return err
}

func BootFailed() cli.Command {
  return cli.Command{
    Name:      "boot-failed",
    Usage:     "Changes the state of the server to BootFailed",
    Action:    BootFailedAction,
    Flags:     []cli.Flag{},
  }
}
