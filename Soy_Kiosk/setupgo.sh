echo "Running Set Up Go"
export PATH=$PATH:/go/bin
export GOPATH=$GOPATH:/go/bin
export GOCACHE=$GOCACHE:/data
go version
echo "build oauth.go"
GOCACHE=/data/ go build OAuth.go
./OAuth
