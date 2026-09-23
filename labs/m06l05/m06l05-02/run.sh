#!/bin/sh
# Concurrent Systems Programming with Go — lesson m06l05 — Shipping in Scratch and Distroless Images
# https://learnsome.tech/courses/go-course/watch?lesson=m06l05
# © LearnSome.tech
set -eu
docker build -t go-course-scratch .
