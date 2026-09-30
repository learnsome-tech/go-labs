#!/bin/sh
set -eu
go version
go build -o probe .
./probe
