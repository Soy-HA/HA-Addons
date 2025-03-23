echo "Running Set Up Go"
export PATH=$PATH:/app/go/bin
export GOCACHE=$GOCACHE:/data
echo "go version" 
go version
echo "go gocache"
go env GOCACHE
echo "Building No Cache"
go build -a OAuth.go
echo "build oauth.go"
GOCACHE=off go build OAuth.go
echo "build main"
GOCACHE=off go build main
echo "build oauth"
GOCACHE=off go build OAuth
./OAuth