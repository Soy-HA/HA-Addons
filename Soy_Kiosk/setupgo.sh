echo "Running Set Up Go"
mkdir "./data/go/"
mkdir "./data/go/cache"
./main
go version
pwd
go env GOCACHE
GOCACHE = /data/go/cache/
export GOCACHE
go env GOCACHE
go build OAuth.go
./data/go/bin/OAuth