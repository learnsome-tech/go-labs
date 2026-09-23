#!/bin/sh
# Concurrent Systems Programming with Go — lesson m01l01 — What Go Is and Why It Matters
# https://learnsome.tech/courses/go-course/watch?lesson=m01l01
# © LearnSome.tech
set -eu
go version
go build -o probe .
./probe
