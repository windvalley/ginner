#!/bin/bash
# change current project name to your own project name.
# e.g. ./change_project_name.sh your-project-name

set -u
set -e


CURRENT_PROJECT_NAME=$(awk 'NR==1{print $2}' go.mod)
NEW_PROJECT_NAME=$1

find . -type f -name "*.go" -o -name "go.mod" |
    xargs gsed -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#"

