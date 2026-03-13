#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/src"
go mod tidy
go build -o ../build/telegram-bot-cli .
