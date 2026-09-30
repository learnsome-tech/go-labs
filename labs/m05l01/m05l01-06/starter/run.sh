#!/bin/sh
set -eu
go run . web01; echo "exit=$?"
