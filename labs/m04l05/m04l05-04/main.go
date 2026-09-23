# Concurrent Systems Programming with Go — lesson m04l05 — Context Cancellation and Timeouts
# https://learnsome.tech/courses/go-course/watch?lesson=m04l05
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("worker stopped")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() { worker(ctx); close(done) }()
	cancel()
	<-done
}
