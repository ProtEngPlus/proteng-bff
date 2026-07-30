#!/usr/bin/env bash
# Regenerates Swagger docs (deterministic + fast, ~0.5s) so docs never drift from the handlers.
if ! command -v swag >/dev/null 2>&1; then
  go install github.com/swaggo/swag/cmd/swag@v1.16.4
fi
swag init >/dev/null

ENV=local go run main.go
