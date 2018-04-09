#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ROOT="${BASH_SOURCE[0]%/*}/.."

cd "$ROOT"

run_unit_test() {
  echo "Running unit tests"
  make test-unit
}

run_style_check() {
  echo "Running style checks"
  make test-lint
}

run_unit_test
run_style_check
