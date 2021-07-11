#!/usr/bin/env bash

set -eu -o pipefail

OUTPUT="$(mktemp)"

cleanup() {
  rm -f "$OUTPUT"
}
trap cleanup EXIT

DESCRIPTION="$1"
DIR="$2"

echo -n "running $DESCRIPTION.. "
cd "$DIR"

shift
shift

set +e
"$@" &> "$OUTPUT"
EXIT_CODE="$?"
set -e

if [[ "$EXIT_CODE" -eq 0 ]]; then
  echo 'success'
  exit 0
fi

echo 'failed'
echo
cat "$OUTPUT"