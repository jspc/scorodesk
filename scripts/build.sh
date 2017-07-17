#!/bin/sh

echo "PATH is: ${PATH}"
echo "GOPATH is: ${GOPATH}"

echo "==== Printing all env vars ===="
printenv
echo "==============================="

make deps

go test -v ./... || exit 1

make build
make upload
