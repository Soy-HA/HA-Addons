echo "Running Set Up Go"
export PATH=$PATH:/app/core/go/bin
export GOCACHE=$GOCACHE:/data
core/go/bin/go version
go env GOCACHE
go build /app/OAuth.go
./core/go/bin/OAuth