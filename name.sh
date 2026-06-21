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
    shift 2
    
    for file in "$@"; do
        if [ -f "$file" ]; then
            if [ "$OS" == "Darwin" ]; then
                sed -i '' "s|$pattern|$replacement|g" "$file"
            else
                sed -i "s|$pattern|$replacement|g" "$file"
            fi
        fi
    done
}

TARGET_FILES=(
    "cmd/api/main.go"
    "README.md"
    "Makefile"
    ".env"
)

do_replace "<org-name>" "$NEW_ORG" "${TARGET_FILES[@]}"
do_replace "<app-name>" "$NEW_APP" "${TARGET_FILES[@]}"