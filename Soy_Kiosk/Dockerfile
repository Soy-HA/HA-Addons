# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
ARG BUILD_FROM
FROM $BUILD_FROM

LABEL \
  io.hass.version="VERSION" \
  io.hass.type="addon" \
  io.hass.arch="armhf|aarch64|i386|amd64"	

WORKDIR /SoyOAuth

RUN apk --no-cache 
RUN mkdir -p /SoyOAuth


# Set the entrypoint script
CMD [""]