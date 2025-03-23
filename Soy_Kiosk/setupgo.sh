echo "Running Set Up Go"
export PATH=$PATH:/app/core/go/bin
go version
go env GOCACHE
go build /app/OAuth.go
./app/core/go/bin/OAuth