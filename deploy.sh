#!/bin/bash

function build() {
  echo "build version ${HASH}"

  go build -o ${BINARY} .
}

function deploy() {
  echo "deploy version ${HASH}"

  scp ${BINARY} home:/usr/local/bin/homeapi
}

ACTION="$1"

HASH=`git rev-parse HEAD`
HASH="${HASH:0:8}"

GOOS="linux"
GOARCH="amd64"

BINARY="homeapi-${GOOS}-${GOARCH}-${HASH}"

case ${ACTION} in
  "build" ) build ;;
  "deploy" ) deploy ;;
  *) echo "Invalid action"; exit 2;
esac
