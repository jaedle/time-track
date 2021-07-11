#!/usr/bin/env bash

set -eu -o pipefail

COMMANDS='curl go lsof mktemp node npm screen'
NODE_VERSION='v14.'
GO_VERSION='go1.16.'

die() {
  local reason="$1"

  echo "preconditions not fulfilled: $reason"
  exit 1
}

require_commands() {
  for cmd in $COMMANDS; do
    which "$cmd" &> /dev/null || die "command '$cmd' not found, please install"
  done
}

require_node_version() {
  node --version | grep "$NODE_VERSION" &> /dev/null || die "please install node with version $NODE_VERSION"
}

require_go_version() {
  go version | grep "$GO_VERSION" &> /dev/null || die "please install node with version $GO_VERSION"
}

require_commands
require_node_version
require_go_version