#!/usr/bin/env bash

if ! hash gometalinter.v1 2>/dev/null ; then
  go get -u gopkg.in/alecthomas/gometalinter.v1
  gometalinter.v1 --install
fi

# Mandatory tests
echo -e "\033[0;31mManadatory Linters: These must pass\033[0m"
gometalinter.v1 --vendor --tests --deadline=20s --disable-all \
--enable=gofmt \
--enable=misspell \
--enable=deadcode \
--enable=ineffassign \
--enable=vet \
./...

mandatory=$?

# Optional tests
echo -e "\033[0;32mOptional Linters: These should pass\033[0m"
gometalinter.v1 --vendor --tests --deadline=20s --disable-all \
--enable=golint \
./...

exit $mandatory
