#!/usr/bin/env bash

REVERSE_PROXY_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

set -eu -o pipefail

"$REVERSE_PROXY_DIR/stop.sh"
screen -dmS reverse-proxy bin/caddy run -config "$REVERSE_PROXY_DIR/Caddyfile"
"$REVERSE_PROXY_DIR/wait-healthy.sh"

