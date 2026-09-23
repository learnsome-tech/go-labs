# Concurrent Systems Programming with Go — lesson m03l02 — Wrapping Errors and errors.Is
# https://learnsome.tech/courses/go-course/watch?lesson=m03l02
# © LearnSome.tech
package main

import (
	"errors"
	"fmt"
)

var ErrBusy = errors.New("busy")

func start() error {
	return fmt.Errorf("start worker: %w", ErrBusy)
}

func main() {
	err := start()
	fmt.Println(err)
	if errors.Is(err, ErrBusy) {
		fmt.Println("retry later")
	}
}
