echo "Running Set Up Go"
export PATH=$PATH:/app/go/bin
#export GOCACHE=$GOCACHE:/data
echo "go version" 
go version
echo "go gocache"
go env GOCACHE
go build -a OAuth.go
./go/bin/OAuth
./data/OAuth