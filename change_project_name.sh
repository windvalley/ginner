#!/bin/bash
# change_project_name.sh
# Change current project name to your own project name.
# e.g. ./change_project_name.sh your-project-name

set -u
set -e


CURRENT_PROJECT_NAME=$(awk 'NR==1{print $2}' go.mod)
NEW_PROJECT_NAME=$1

# shellcheck disable=SC2038
find . -type f -name "*.go" -o -name "go.mod" |
    xargs gsed -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#"

gsed -i "/^BIN_NAME=/s/${CURRENT_PROJECT_NAME}/${NEW_PROJECT_NAME}/" ./service.sh
gsed -i "s/${CURRENT_PROJECT_NAME}/${NEW_PROJECT_NAME}/g" \
    ./Dockerfile ./conf/*.toml .gitignore supervisord.conf \
    "${CURRENT_PROJECT_NAME}".service

mv "${CURRENT_PROJECT_NAME}".service "${NEW_PROJECT_NAME}".service

