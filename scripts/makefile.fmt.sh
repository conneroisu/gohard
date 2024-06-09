#!/bin/bash

gofmt -w .

golines -w --max-len=79 .
