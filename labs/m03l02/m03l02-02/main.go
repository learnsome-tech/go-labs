# Concurrent Systems Programming with Go — lesson m03l02 — Wrapping Errors and errors.Is
# https://learnsome.tech/courses/go-course/watch?lesson=m03l02
# © LearnSome.tech
package main

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("missing")

func load() error {
	return fmt.Errorf("read config: %w", ErrMissing)
}

func main() {
	err := load()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrMissing))
}
