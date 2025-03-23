echo "Running Set Up Go"
tar -xzf "go1.24.1.src.tar.gz" /app/core/
export PATH=$PATH:/app/core/go/bin
go version
go env GOCACHE
go build /app/OAuth.go
./app/core/go/bin/OAuth