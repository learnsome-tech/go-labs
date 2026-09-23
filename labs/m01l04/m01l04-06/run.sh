#!/bin/sh
# Concurrent Systems Programming with Go — lesson m01l04 — Your First Program and go run
# https://learnsome.tech/courses/go-course/watch?lesson=m01l04
# © LearnSome.tech
set -eu
go build -o diskcheck . && ./diskcheck
