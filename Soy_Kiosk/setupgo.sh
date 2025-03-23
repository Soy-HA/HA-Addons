echo "Running Set Up Go"
mkdir "./data/go/"
mkdir "./data/go/cache"
./go/out/main
go version
go env GOCACHE
go build OAuth.go
./data/go/bin/OAuth