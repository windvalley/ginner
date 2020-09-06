#!/usr/bin/env bash
# code_quality_inspect.sh


# Markdown
# install:
#    npm install -g markdownlint-cli
markdownlint ./*.md

# Dockerfile
# install:
#    brew install hadolint
hadolint ./Dockerfile

# Shell
# install:
#    brew install shellcheck
shellcheck ./*.sh

# Go
# install:
#    go get -u -v github.com/mgechev/revive
revive ./...

# Go
# install:
#    brew tap alecthomas/homebrew-tap
#    brew install gometalinter
#    go get github.com/gojp/goreportcard/cmd/goreportcard-cli
goreportcard-cli -v
