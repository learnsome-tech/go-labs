#!/bin/sh
set -eu
go run .
go build -o diskcheck .
./diskcheck
go run . more arguments
