ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
#Install GO 
WORKDIR /app/ 
RUN curl -LO "https://go.dev/dl/go1.24.1.linux-amd64.tar.gz"
RUN tar -xzvf "go1.24.1.linux-amd64.tar.gz"
RUN rm "go1.24.1.linux-amd64.tar.gz"
# Copy data for add-on
COPY setupgo.sh OAuth.go /app/
RUN chmod a+x /app/setupgo.sh

RUN export PATH=$PATH:/app/go/bin
RUN export GOPATH=$GOPATH:/app/go/bin
RUN export GOCACHE=$GOCACHE:/data
RUN go version
RUN GOCACHE=/data/ go build -tags debug OAuth.go



CMD [ "./OAuth" ]