#!/bin/sh
set -eu
go build -o diskcheck . && ./diskcheck
