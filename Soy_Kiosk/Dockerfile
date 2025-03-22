# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
ARG BUILD_FROM
FROM $BUILD_FROM

COPY run.sh
RUN chmod a+x /run.sh
RUN mkdir -p /SoyOAuth

# Set the entrypoint script
CMD ["/run.sh"]