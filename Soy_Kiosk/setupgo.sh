echo "Running Set Up Go"
mkdir "/data/go/cache"
go version
go build OAuth.go
ls
echo "dumping go bin"
ls /usr/local/go/
ls /data/go/bin/
./data/OAuth