#!/bin/bash

NEW_ORG=$1
NEW_APP=$2

if [ -z "$NEW_ORG" ] || [ -z "$NEW_APP" ]; then
    echo "Usage: ./name.sh <org-name> <app-name>"
    exit 1
fi

OS="$(uname)"

do_replace() {
    local pattern=$1
    local replacement=$2
    if [ "$OS" == "Darwin" ]; then
        sed -i '' "s|$pattern|$replacement|g" cmd/api/main.go
        sed -i '' "s|$pattern|$replacement|g" README.md
        sed -i '' "s|$pattern|$replacement|g" Makefile
        sed -i '' "s|$pattern|$replacement|g" .env
    else
        sed -i "s|$pattern|$replacement|g" cmd/api/main.go
        sed -i "s|$pattern|$replacement|g" README.md
        sed -i "s|$pattern|$replacement|g" Makefile
        sed -i "s|$pattern|$replacement|g" .env
    fi
}


do_replace "<org-name>" "$NEW_ORG"
do_replace "<app-name>" "$NEW_APP"
