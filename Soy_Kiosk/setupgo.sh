echo "Running Set Up Go"
export PATH=$PATH:/app/go/bin
#export GOCACHE=$GOCACHE:/data
echo "go version" 
go version
echo "go gocache"
go env GOCACHE
echo "Building No Cache"
go build -a OAuth.go
./OAuth