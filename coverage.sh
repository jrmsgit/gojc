#!/bin/sh -eu
ARGS=${@:-'./...'}
go test -coverprofile coverage.prof ${ARGS}
go tool cover -html coverage.prof -o coverage.html
