package api

import (
  "net/http"
)

type HTML5Dir struct {
  Dir     http.FileSystem
  Default string
}

func (d HTML5Dir) Open(name string) (http.File, error) {
  f, err := d.Dir.Open(name)
  if err != nil {
    f, err = d.Dir.Open(d.Default)
    if err != nil {
      return nil, err
    }
  }
  return f, nil
}
