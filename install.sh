#!/bin/sh -eu
ARGS=${@:-'./...'}
go env | grep -E 'GOBIN|GOEXE|GOPATH' | sort
go install -i ${ARGS}
