#!/usr/bin/env bash
# code_quality_inspect.sh


# Markdown
# install markdownlint-cli:
#    npm install -g markdownlint-cli
#markdownlint ./*.md

# Markdown
# install mdl:
#    sudo gem install mdl
MDL_IGNORE_RULES="MD013"
MDL_APPLY_RULES=$(mdl -l | grep -Eo '^MD[0-9]{3}' | grep -Ev $MDL_IGNORE_RULES |
    xargs | sed 's/ /,/g')

mdl -r "$MDL_APPLY_RULES" ./*.md

# Dockerfile
# install hadolint:
#    brew install hadolint
hadolint ./Dockerfile

# Shell
# install shellcheck:
#    brew install shellcheck
shellcheck ./*.sh

# Go
# install revive:
#    go get -u -v github.com/mgechev/revive
revive ./...

# Go
# install golangci-lint:
#    go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0
golangci-lint run

# Go
# install goreportcard-cli:
#    brew tap alecthomas/homebrew-tap
#    brew install gometalinter
#    go get github.com/gojp/goreportcard/cmd/goreportcard-cli
goreportcard-cli -v
