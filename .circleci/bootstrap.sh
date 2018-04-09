#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ROOT="${BASH_SOURCE[0]%/*}/.."

cd "$ROOT"

make bootstrap
