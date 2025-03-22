# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
ARG BUILD_FROM
FROM $BUILD_FROM

WORKDIR /SoyOAuth

RUN apk --no-cache add gcc musl-dev git sqlite-dev curl jq
RUN mkdir -p /SoyOAuth


# Set the entrypoint script
CMD ["echo 'hi'"]