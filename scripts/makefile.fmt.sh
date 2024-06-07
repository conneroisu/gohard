#!/bin/bash
# file: makefile.fmt.sh
# url: https://github.com/conneroisu/hardgo/scripts/makefile.fmt.sh
# title: Formatting Go Files
# description: This script formats the Go files using gofmt and golines.
#
# Usage: make fmt

gofmt -w .

goline -w --max-len=79 .
