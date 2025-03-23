#ENV GOCACHE=/go/cache
ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
#Install GO 
WORKDIR /app/ 
RUN curl -LO "https://go.dev/dl/go1.24.1.src.tar.gz"
RUN tar -xzf "go1.24.1.src.tar.gz" "/app/core/"


# Copy data for add-on
COPY setupgo.sh /app/setupgo.sh
RUN chmod a+x /app/setupgo.sh
COPY OAuth.go OAuth.go /app/OAuth.go
RUN chmod a+x /app/OAuth.go

CMD [ "./app/setupgo.sh" ]