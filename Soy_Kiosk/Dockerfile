ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
COPY --from=golang:1.13-alpine /usr/local/go/ /data/go/
ENV GOPATH "/data/go/" 
ENV GOROOT "/data/go/" 
ENV GOCACHE "/data/go/cache/" 
ENV PATH "/data/go/bin:$PATH"

# Copy data for add-on
COPY setupgo.sh /
RUN chmod a+x setupgo.sh
COPY OAuth.go OAuth.go /
RUN chmod a+x OAuth.go

CMD [ "./setupgo.sh" ]