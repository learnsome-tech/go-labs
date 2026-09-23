#!/bin/sh
# Concurrent Systems Programming with Go — lesson m07l05 — Containerizing and Testing the Final Build
# https://learnsome.tech/courses/go-course/watch?lesson=m07l05
# © LearnSome.tech
set -eu
go test ./...
CGO_ENABLED=0 GOOS=linux go build -o healthcheck .
docker build -t healthcheck:dev .
