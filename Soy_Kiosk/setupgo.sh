echo "Running Set Up Go"
export PATH=$PATH:/app/go/bin
export GOCACHE=$GOCACHE:/data
go/bin/go version
go env GOCACHE
go build /app/OAuth.go
./go/bin/OAuth
./data/OAuth