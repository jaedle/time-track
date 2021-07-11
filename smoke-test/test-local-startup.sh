#!/usr/bin/env bash

set -eu -o pipefail

die() {
  reason="$1"

  echo "$reason"
  exit 1
}

test_reverse_proxy() {
  curl -fsSL http://localhost:3000/health || die 'reverse proxy is not healthy'
}

test_ui() {
  curl -fsSL http://localhost:3000/ || die 'ui is not healthy'
}

test_service() {
  curl -fsSL http://localhost:3000/api/health || die 'service is not healthy'

}

test_reverse_proxy
test_ui
test_service