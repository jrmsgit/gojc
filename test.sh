#!/bin/sh -eu
ARGS=${@:-'./...'}
extra=${GOJC_TEST:-''}
if test "${extra}" = 'race'; then
  ARGS="-race ${ARGS}"
fi
if test "${extra}" = 'bench'; then
  ARGS="-bench=. ${ARGS}"
fi
echo "go test ${ARGS}"
go test ${ARGS}
