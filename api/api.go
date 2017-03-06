package api

import (
  "fmt"
  "github.com/griff/thonix-backend/server"
  "github.com/gorilla/mux"
  apiGraphql "github.com/griff/thonix-backend/api/graphql"
  graphql "github.com/neelance/graphql-go"
  "github.com/neelance/graphql-go/relay"
  "net/http"
)

type API struct {
  *server.Server
  Assets http.FileSystem
  Shell string
}

func (api API) Router() (*mux.Router, error) {
  sh := api.ShellHandler()

  relay, err := api.RelayHandler()
  if err != nil {
    return nil, err
  }

  r := mux.NewRouter()
  r.HandleFunc("/terminals/{pid}", sh.ShellWSHandler)
  r.HandleFunc("/terminals", sh.MakeProcess).Methods("POST")
  r.Handle("/graphql", relay)
  r.HandleFunc("/state", api.ServerState).Methods("GET")
  r.PathPrefix("/").Handler(http.FileServer(HTML5Dir{
    Dir:     api.Assets,
    Default: "index.html",
  }))
  return r, nil
}

func (api API) Schema() (*graphql.Schema, error) {
  return apiGraphql.Schema(api.Server)
}

func (api API) ShellHandler() *ShellHandler {
  return NewShellHandler(api.Shell)
}

func (api API) RelayHandler() (*relay.Handler, error) {
  schema, err := api.Schema()
  if err != nil {
    return nil, err
  }
  return &relay.Handler{Schema: schema}, nil
}

func (api API) ServerState(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  fmt.Fprintf(w, "%s", api.Server.Status().String())
}