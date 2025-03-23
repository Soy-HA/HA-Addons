echo "Running Set Up Go"
mkdir "./data/go/"
mkdir "./data/go/cache"
go version
export GOCACHE = "/data/go/cache/"
go build OAuth.go
./data/go/bin/OAuth