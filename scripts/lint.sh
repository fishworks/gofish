#!/usr/bin/env bash

if ! hash golangci-lint 2>/dev/null ; then
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
fi

# mandatory tests
golangci-lint \
  --tests \
  --disable-all \
  --enable=gofmt \
  --enable=misspell \
  --enable=deadcode \
  --enable=ineffassign \
  --enable=govet \
  run ./...

mandatory=$?

# optional tests
golangci-lint \
  --tests \
  --disable-all \
  --enable=golint \
  run ./...

exit $mandatory
