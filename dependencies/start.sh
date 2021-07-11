#!/usr/bin/env bash

DEPENDENCIES_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

set -eu -o pipefail

docker-compose up -d
"$DEPENDENCIES_DIR/wait-mariadb-healthy.sh"

