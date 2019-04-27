#!/bin/sh -eu
go vet ./...
gofmt -s -l .
exec ./test.sh
