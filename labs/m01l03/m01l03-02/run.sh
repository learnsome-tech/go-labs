#!/bin/sh
# Concurrent Systems Programming with Go — lesson m01l03 — Go Workspaces and Modules
# https://learnsome.tech/courses/go-course/watch?lesson=m01l03
# © LearnSome.tech
set -eu
mkdir probe && cd probe
go mod init course.local/probe
cat go.mod
