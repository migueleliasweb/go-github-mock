#!/usr/bin/env bash

if [ "$(git branch --show-current)" != "master" ]; then
    echo "Not in master branch. Aborting."
    exit 1
fi

BRANCH_NAME="update-gh-definitions-$(date '+%s')"

git checkout -b ${BRANCH_NAME}

go run main.go

if [ "$(git status --porcelain | wc -l)" != "1" ]; then
    echo "Found more changes than expected. Aborting."
    exit 1
fi

git commit -am "Update GH definitions"

git push origin ${BRANCH_NAME}