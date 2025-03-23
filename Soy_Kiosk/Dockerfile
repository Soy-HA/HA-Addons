ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
ENV GOPATH "/usr/local/go/" 
ENV GOROOT "/usr/local/go/" 
ENV GOCACHE "/usr/local/go/cache/" 
ENV HOME "/usr/local/go/cache/" 
ENV PATH "/usr/local/go//bin:$PATH"

# Python 3 HTTP Server serves the current working dir
# So let's set it to our add-on persistent data directory.
#WORKDIR /data

# Copy data for add-on
COPY setupgo.sh /
RUN chmod a+x setupgo.sh
COPY OAuth.go OAuth.go /
RUN chmod a+x OAuth.go

CMD [ "./setupgo.sh" ]