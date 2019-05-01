#!/bin/sh -eu
ARGS=${@:-'./...'}
gofmt -w -s -l .
go vet ${ARGS}
exec ./test.sh ${ARGS}
