#!/usr/bin/env bash

set -eu -o pipefail

REVERSE_PROXY_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

install() {
  go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
}

clean() {
  rm -rf "${REVERSE_PROXY_DIR:?}/bin"
}

build() {
  mkdir -p "$REVERSE_PROXY_DIR/bin"
  xcaddy build \
    --output "$REVERSE_PROXY_DIR/bin/caddy"
}

install
clean
build