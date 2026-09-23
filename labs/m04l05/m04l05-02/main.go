# Concurrent Systems Programming with Go — lesson m04l05 — Context Cancellation and Timeouts
# https://learnsome.tech/courses/go-course/watch?lesson=m04l05
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	fmt.Println(work(ctx))
}
