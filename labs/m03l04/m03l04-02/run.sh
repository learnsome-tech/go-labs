#!/bin/sh
# Concurrent Systems Programming with Go — lesson m03l04 — The Standard testing Package
# https://learnsome.tech/courses/go-course/watch?lesson=m03l04
# © LearnSome.tech
set -eu
go test -count=1 ./...
