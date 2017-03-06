package assets

import (
  "github.com/elazarl/go-bindata-assetfs"
  "net/http"
)

func AssetFS() http.FileSystem {
  return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "frontend"}
}