echo "Running Set Up Go"
mkdir "./data/go/"
mkdir "./data/go/cache"
go version
pwd
go env GOCACHE
GOCACHE = /data/go/cache/
export GOCACHE
go enc GOCACHE
go build OAuth.go
./data/go/bin/OAuth