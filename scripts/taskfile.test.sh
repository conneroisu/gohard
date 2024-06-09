#!/bin/bash
# url: https://github.com/conneroisu/hardgo/blob/main/scripts/taskfile.test.sh

go test -race -v -timeout 30s ./...

go test -coverprofile=coverage.out ./...

# if gocovsh is executable
if [ -x "$(command -v gocovsh)" ]; then
    # if gocovsh is not empty
    if [ -s coverage.out ]; then
        # run gocovsh
        gocovsh
    else
        # if coverage.out is empty/not found
        echo "No coverage.out file found."
    fi
else
    # if gocovsh is not executable
    echo "gocovsh is not executable."
fi
