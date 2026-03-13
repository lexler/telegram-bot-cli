#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
mkdir -p build
go mod tidy
go build -o build/telegram-bot-cli .
