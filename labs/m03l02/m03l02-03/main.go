# Concurrent Systems Programming with Go — lesson m03l02 — Wrapping Errors and errors.Is
# https://learnsome.tech/courses/go-course/watch?lesson=m03l02
# © LearnSome.tech
package main

import (
	"errors"
	"fmt
)

var ErrMissing = errors.New("missing")

func main() {
	err := fmt.Errorf("open settings: %w", ErrMissing)
	if errors.Is(err, ErrMissing) {
		fmt.Println("use defaults")
		return
	}
	fmt.Println("stop")
}
