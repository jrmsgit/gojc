#!/bin/sh -eu
ARGS=${@:-'./...'}
go test ${ARGS}
