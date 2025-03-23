echo "Running Set Up Go"
mkdir "./data/go/"
mkdir "./data/go/cache"
go version
go build OAuth.go
echo "dump 1"
ls
echo "dump 2"
ls ./data/
echo "dump 3"
ls ./data/go/
echo "dump 4"
ls ./usr/local/go/
echo "dump 5"
ls ./data/go/bin/
export GOCACHE = ./data/go/cache/
./data/go/bin/OAuth