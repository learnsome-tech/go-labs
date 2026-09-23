#!/bin/sh
# Concurrent Systems Programming with Go — lesson m03l05 — Writing Table Driven Tests
# https://learnsome.tech/courses/go-course/watch?lesson=m03l05
# © LearnSome.tech
set -eu
go test -count=1 ./...
