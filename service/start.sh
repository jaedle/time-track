#!/usr/bin/env bash

SERVICE_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

set -eu -o pipefail

"$SERVICE_DIR/stop.sh"
screen -dmS service go run main.go
"$SERVICE_DIR/wait-healthy.sh"

