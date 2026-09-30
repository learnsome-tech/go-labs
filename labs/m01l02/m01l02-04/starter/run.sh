#!/bin/sh
set -eu
gofmt -l .
gofmt -w main.go
gofmt -l .
