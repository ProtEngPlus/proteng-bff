# proteng-bff: API gateway between protengplus-frontend and the backend services.
# Run `make` to list targets. Works from Git Bash, cmd or PowerShell (needs Git for Windows).

# On Windows, `bash` found from cmd or PowerShell is often the WSL launcher
# (C:\Windows\System32\bash.exe), so use the bash that ships with Git for Windows.
# Git's layout is <root>/mingw64/libexec/git-core and <root>/bin/bash.exe.
ifeq ($(OS),Windows_NT)
GIT_EXEC_PATH := $(shell git --exec-path)
ifneq ($(findstring /mingw64/libexec/git-core,$(GIT_EXEC_PATH)),)
SHELL := $(subst /mingw64/libexec/git-core,/bin/bash.exe,$(GIT_EXEC_PATH))
else
SHELL := bash
endif
else
SHELL := bash
endif
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help
MAKEFLAGS += --no-print-directory

# python3 on Windows is often the Microsoft Store stub, so check that it runs.
PYTHON ?= $(shell python3 -c 'import sys' >/dev/null 2>&1 && echo python3 || echo python)
IMAGE ?= proteng-bff
SWAG_VERSION ?= v1.16.4

.PHONY: help setup hooks env deps run swagger fmt lint test check build docker-build

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*## "} /^##@/ {printf "\n%s\n", substr($$0, 5)} /^[a-zA-Z0-9_.-]+:.*## / {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

##@ Setup
setup: hooks env deps ## First-time setup: git hooks, .env.local and Go modules (safe to re-run)

hooks: ## Install the git hooks (needs `pip install pre-commit`)
	$(PYTHON) -m pre_commit install --hook-type pre-commit --hook-type pre-push --hook-type commit-msg

env: ## Create .env.local from .env.example (never overwrites)
	@if [ -f .env.local ]; then echo ".env.local exists, left as is"; else cp .env.example .env.local; echo "created .env.local from .env.example"; fi

deps: ## Download the Go modules
	go mod download

##@ Run
run: ## Regenerate Swagger, then run with ENV=local (reads .env.local)
	bash run.sh

swagger: ## Regenerate docs/ from the handler annotations (installs swag if missing)
	@command -v swag >/dev/null 2>&1 || go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)
	swag init

##@ Checks
fmt: ## Format every Go file with gofmt
	gofmt -l -w .

lint: ## Fail on files gofmt would change, then go vet
	@out="$$(gofmt -l .)"; if [ -n "$$out" ]; then echo "not gofmt'd (run: make fmt):"; echo "$$out"; exit 1; fi
	go vet ./...

test: ## Run the unit tests
	go test ./...

check: lint test ## Everything CI checks

build: ## Compile every package
	go build ./...

docker-build: ## Build the image locally
	docker build -t $(IMAGE) .
