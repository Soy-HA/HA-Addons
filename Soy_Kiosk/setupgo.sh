echo "Running Set Up Go"
export PATH=$PATH:/app/go/bin
export GOCACHE=$GOCACHE:/data
go version
echo "build oauth.go"
GOCACHE=/data/ go build OAuth.go
./OAuth