package rpc

import (
  "golang.org/x/net/context"
  "net"
  capnp "zombiezen.com/go/capnproto2"
  "zombiezen.com/go/capnproto2/rpc"
)

type RCPClient struct {
  conn net.Conn
  c *rpc.Conn
  cancel context.CancelFunc
  rpc Server
  ctx context.Context
}

func ConnectPublic(ctx context.Context, socketPath string) (*RCPClient, error) {
  if socketPath == "" {
    socketPath = Server_socketPath
  }
  ctx, cancel := context.WithCancel(ctx)
  conn, err := net.Dial("unix", socketPath)
  if err != nil {
    return nil, err
  }
  c := rpc.NewConn(rpc.StreamTransport(conn))
  return &RCPClient {
    conn: conn,
    c: c,
    ctx: ctx,
    cancel: cancel,
    rpc: Server {Client: c.Bootstrap(ctx)},
  }, nil
}

func (c *RCPClient) Close() error {
  c.cancel()
  return c.conn.Close()
}

func (c *RCPClient) State(params func(Server_state_Params) error, opts ...capnp.CallOption) Server_state_Results_Promise {
  return c.rpc.State(c.ctx, params, opts...)
}

type RCPAdminClient struct {
  conn net.Conn
  c *rpc.Conn
  cancel context.CancelFunc
  ctx context.Context
  rpc ServerAdmin
}

func ConnectAdmin(ctx context.Context, socketPath string) (*RCPAdminClient, error) {
  if socketPath == "" {
    socketPath = ServerAdmin_socketPath
  }
  ctx, cancel := context.WithCancel(ctx)
  conn, err := net.Dial("unix", socketPath)
  if err != nil {
    return nil, err
  }
  c := rpc.NewConn(rpc.StreamTransport(conn))
  return &RCPAdminClient {
    conn: conn,
    c: c,
    ctx: ctx,
    cancel: cancel,
    rpc: ServerAdmin {Client: c.Bootstrap(ctx)},
  }, nil
}

func (c *RCPAdminClient) Close() error {
  c.cancel()
  return c.conn.Close()
}

func (c *RCPAdminClient) State(params func(Server_state_Params) error, opts ...capnp.CallOption) Server_state_Results_Promise {
  return c.rpc.State(c.ctx, params, opts...)
}

func (c *RCPAdminClient) SetState(params func(ServerAdmin_setState_Params) error, opts ...capnp.CallOption) ServerAdmin_setState_Results_Promise {
  return c.rpc.SetState(c.ctx, params, opts...)
}
