#!/bin/bash
# go get -u github.com/mitchellh/gox


OSARCH="linux/amd64 darwin/amd64"

gox -osarch="$OSARCH" || {
    go get -u github.com/mitchellh/gox
    gox -osarch="$OSARCH"
}

