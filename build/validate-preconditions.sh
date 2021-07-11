#!/usr/bin/env bash

set -eu -o pipefail

COMMANDS='npm node'
NODE_VERSION='v14.'

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

require_node() {
  node --version | grep "$NODE_VERSION" &> /dev/null || die "please install node with version $NODE_VERSION"
}


require_commands
require_node