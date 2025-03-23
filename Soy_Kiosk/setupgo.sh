echo "Running Set Up Go"
ls
echo "dumping /app/"
ls /app/
cd app/
mkdir core/
cd ../
tar -xzf "go1.24.1.darwin-amd64.tar.gz" /app/core/
export PATH=$PATH:/app/core/go/bin
ls /app/core/go/
/app/core/go/bin/go version
go env GOCACHE
go build /app/OAuth.go
./app/core/go/bin/OAuth