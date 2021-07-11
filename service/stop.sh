#!/usr/bin/env bash

set -eu -o pipefail

if lsof -ti tcp:8080 &> /dev/null; then
  # shellcheck disable=SC2046
  kill -9 $(lsof -ti tcp:8080) &> /dev/null
fi