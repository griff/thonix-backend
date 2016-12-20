
GOPATH:=${PWD}:$(value GOPATH)
.PHONE: frontend-dist

bindata_assetfs.go: frontend-dist
	go-bindata-assetfs -prefix ../frontend/release ../frontend/release/...

build: vendor bindata_assetfs.go
	go build -v github.com/griff/thonix
frontend-dist:
	$(MAKE) -C ../frontend dist
vendor: glide.lock glide.yaml
	glide install

