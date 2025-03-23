echo "Running Set Up Go"
mkdir "/usr/local/go/cache/"
go version
go build /OAuth.go
ls
echo "dumping go bin"
echo $GOPATH
./OAuth