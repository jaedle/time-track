#!/usr/bin/env bash

set -eu -o pipefail

NODE_VERSION='v14.'

die() {
  local reason="$1"

  echo "preconditions not fulfilled: $reason"
  exit 1
}

require_node() {
  node --version | grep "$NODE_VERSION" &> /dev/null || die "please install node with version $NODE_VERSION"
}

require_node