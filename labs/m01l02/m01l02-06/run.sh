#!/bin/sh
# Concurrent Systems Programming with Go — lesson m01l02 — Installing the Go Toolchain
# https://learnsome.tech/courses/go-course/watch?lesson=m01l02
# © LearnSome.tech
set -eu
go doc fmt.Println | sed -n '1,3p'
go doc os.Getenv | sed -n '3,4p'
