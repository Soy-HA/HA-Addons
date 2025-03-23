echo "Running Set Up Go"
mkdir core
tar -xzf "go1.24.1.darwin-amd64.tar.gz" core/
export PATH=$PATH:/app/core/go/bin
ls core/go/
core/go/bin/go version
go env GOCACHE
go build /app/OAuth.go
./core/go/bin/OAuth