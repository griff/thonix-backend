
GOPATH:=${PWD}:$(value GOPATH)

build: vendor
	go build -v github.com/griff/thonix

bindata_assetfs.go: frontend-dist
	go-bindata-assetfs -prefix ../frontend/release ../frontend/release/...

frontend-dist:
	$(MAKE) -C ../frontend dist
vendor: glide.lock glide.yaml
	glide install

.DEFAULT: build
.PHONY: frontend-dist
