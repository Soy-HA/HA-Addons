ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
ENV GOPATH "/usr/local/go/" 
ENV GOROOT "/usr/local/go/" 
ENV GOCACHE "/data/go/cache/" 
ENV PATH "/usr/local/go/bin:$PATH"

# Copy data for add-on
COPY setupgo.sh /
RUN chmod a+x setupgo.sh
COPY OAuth.go OAuth.go /
RUN chmod a+x OAuth.go

CMD [ "./setupgo.sh" ]