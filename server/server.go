package server;

import (
  "sync"
  "github.com/griff/thonix/rpc"
)

type Server struct {
  bootSteps int32
  status rpc.Status
  mu   sync.Mutex
}

func NewServer(bootSteps int32) (*Server, error) {
  return &Server {
    bootSteps: bootSteps,
    status: rpc.Status_booting,
  }, nil
}

func (s *Server) Boot() (*Boot, error) {
  return NewBoot(s.bootSteps)
}

func (s *Server) Status() rpc.Status {
  s.mu.Lock()
  status := s.status
  s.mu.Unlock()
  return status
}

func (s *Server) SetStatus(status rpc.Status) {
  s.mu.Lock()
  s.status = status
  s.mu.Unlock()
}

func (server *Server) SetStatusFromString(status string) error {
  s, err := rpc.ValidStatusFromString(status)
  if err != nil {
    return err
  }
  server.SetStatus(s)
  return nil
}