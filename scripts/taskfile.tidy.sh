#!/bin/bash
# url: https://github.com/conneroisu/hardgo/scripts/makefile.tidy.sh

gum spin --spinner dot --title "Running Go Mod Tidy" --show-output -- \
    go mod tidy
