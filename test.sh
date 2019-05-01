#!/bin/sh -eu
ARGS=${@:-'./...'}
extra=${GOJC_TEST:-''}
if test "X${extra}" != 'X'; then
  ARGS="-${extra} ${ARGS}"
fi
go test ${ARGS}
