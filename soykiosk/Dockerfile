# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
ARG BUILD_FROM
FROM $BUILD_FROM

LABEL \
  io.hass.version="1.0.0" \
  io.hass.type="addon" \
  io.hass.arch="armhf|aarch64|i386|amd64"	

#WORKDIR /SoyOAuth

COPY run.sh
RUN chmod a+x /run.sh
RUN mkdir -p /SoyOAuth


# Set the entrypoint script
CMD ["/run.sh"]