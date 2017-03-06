package server

import (
  "github.com/docker/docker/pkg/listeners"
  thonixrpc "github.com/griff/thonix-backend/rpc"
  "github.com/griff/thonix-backend/server"
  "net"
  capnp "zombiezen.com/go/capnproto2"
  "zombiezen.com/go/capnproto2/rpc"
)

type RPCServer struct {
  *server.Server
  public net.Listener
  admin net.Listener
}

func NewServer(s *server.Server) (*RPCServer, error) {
  return &RPCServer{Server: s}, nil
}

func (rs *RPCServer) State(call thonixrpc.Server_state) error {
  call.Results.SetState(rs.Server.Status())
  return nil
}

func (rs *RPCServer) SetState(call thonixrpc.ServerAdmin_setState) error {
  state := call.Params.NewState()
  err := state.Valid()
  if err != nil {
    return err
  }
  rs.Server.SetStatus(state)
  return nil
}

func (rs *RPCServer) ListenPublic(socketPath string) error {
    main := thonixrpc.Server_ServerToClient(rs)

    if socketPath == "" {
      socketPath = thonixrpc.Server_socketPath
    }
    ls, err := listeners.Init("unix", socketPath, "users", nil)
    if err != nil {
      return err
    }
    rs.public = ls[0]
    if err != nil {
      return err
    }
    go handleServer(rs.public, main.Client)
    return nil
}

func (rs *RPCServer) ListenAdmin(socketPath string) error {
    main := thonixrpc.ServerAdmin_ServerToClient(rs)

    if socketPath == "" {
      socketPath = thonixrpc.ServerAdmin_socketPath
    }
    ls, err := listeners.Init("unix", socketPath, "adm", nil)
    if err != nil {
      return err
    }
    rs.admin = ls[0]
    go handleServer(rs.admin, main.Client)
    return nil
}

func (rs *RPCServer) Close() error {
  var err, err2 error
  if rs.public != nil {
    err = rs.public.Close()
  }
  if rs.admin != nil {
    err2 = rs.admin.Close()
  }
  if err != nil {
    return err
  }
  return err2
}

func handleServer(l net.Listener, client capnp.Client) error {
  var err error
  for {
    var c net.Conn
    c, err = l.Accept()
    if err != nil {
      break
    }
    go handleServerConn(c, client)
  }
  l.Close()
  return err
}

func handleServerConn(c net.Conn, client capnp.Client) error {
    // Listen for calls, using the HashFactory as the bootstrap interface.
    conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(client))
    // Wait for connection to abort.
    err := conn.Wait()
    return err
}