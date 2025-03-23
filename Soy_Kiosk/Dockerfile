ARG BUILD_FROM
FROM $BUILD_FROM

# Install requirements for add-on
RUN \
  apk add --no-cache
  
#Install GO 
WORKDIR /app/ 
RUN curl -LO "https://go.dev/dl/go1.24.1.src.tar.gz"

# Copy data for add-on
COPY setupgo.sh OAuth.go /app/
RUN chmod a+x /app/setupgo.sh

CMD [ "./app/setupgo.sh" ]