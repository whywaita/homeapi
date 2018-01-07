#!/bin/bash

function build() {
  echo "build version ${HASH}"

  GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BINARY} .
}

function deploy() {
  echo "deploy version ${HASH}"

  scp ${BINARY} home:yayoi
}

function production() {
  export GOOS="linux"
  export GOARCH="amd64"
}

function develop() {
  export GOOS="darwin"
  export GOARCH="amd64"
}

STAGE="$1"
ACTION="$2"

HASH=`git rev-parse HEAD`
HASH="${HASH:0:8}"

case ${STAGE} in
  "develop" ) develop ;;
  "prod" ) production ;;
  *) echo "Invalid stage"; exit 2;
esac

BINARY="yayoi-${GOOS}-${GOARCH}-${HASH}"

case ${ACTION} in
  "build" ) build ;;
  "deploy" ) deploy ;;
  *) echo "Invalid action"; exit 2;
esac
