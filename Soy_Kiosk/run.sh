#!/usr/bin/with-contenv bashio
VERSION="$(bashio::config 'version')"
echo "running run"
echo "Starting a new instance $version"

./setupgo.sh