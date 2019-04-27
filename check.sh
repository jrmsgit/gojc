#!/bin/sh -eu
ARGS=${@:-'./...'}
go vet ${ARGS}
gofmt -s -l .
exec ./test.sh ${ARGS}
