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

START_TIME="$SECONDS"
set +e
"$@" &> "$OUTPUT"
EXIT_CODE="$?"
set -e

ELAPSED_TIME=$((SECONDS - START_TIME))

if [[ "$EXIT_CODE" -eq 0 ]]; then
  echo "success [${ELAPSED_TIME}s]"
  exit 0
fi

echo 'failed'
echo
cat "$OUTPUT"
exit "$EXIT_CODE"