#!/bin/sh -eu
ARGS=${@:-'./...'}
go vet ${ARGS}
gofmt -w -s -l .
exec ./test.sh ${ARGS}
