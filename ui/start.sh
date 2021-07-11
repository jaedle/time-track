#!/usr/bin/env bash

UI_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

set -eu -o pipefail

cd "$UI_DIR/"

./stop.sh
screen -dmS ui npm run start
./wait-healthy.sh
