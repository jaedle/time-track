#!/usr/bin/env bash

set -eu -o pipefail

REPOSITORY="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )/../"

cd "$REPOSITORY"

COMMANDS='curl go lsof mktemp node npm screen'
NODE_VERSION='v14.'
GO_VERSION='go1.16.'
PROJECT_COMMANDS='tools/bin/golangci-lint'

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

require_project_commands() {
  for cmd in $PROJECT_COMMANDS; do
    which "$cmd" &> /dev/null || die "command '$cmd' not found, please run task install"
  done
}

require_node_version() {
  node --version | grep "$NODE_VERSION" &> /dev/null || die "please install node with version $NODE_VERSION"
}

require_go_version() {
  go version | grep "$GO_VERSION" &> /dev/null || die "please install node with version $GO_VERSION"
}

require_commands
require_project_commands
require_node_version
require_go_version