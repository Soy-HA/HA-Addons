FROM golang:latest as builder 
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /out/main .
ENV GOCACHE=/go/cache
RUN go env GOCACHE # Verify the GOCACHE path during build

WORKDIR /root/


ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
#COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
#ENV GOPATH "/usr/local/go/" 
#ENV GOROOT "/usr/local/go/" 
#ENV GOCACHE "/data/go/cache/" 
#ENV PATH "/usr/local/go/bin:$PATH"

# Copy data for add-on
COPY setupgo.sh /
RUN chmod a+x setupgo.sh
COPY OAuth.go OAuth.go /
RUN chmod a+x OAuth.go

 COPY --from=builder /out/main .
CMD [ "./setupgo.sh" ]