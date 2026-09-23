# Concurrent Systems Programming with Go — lesson m04l04 — The context Package
# https://learnsome.tech/courses/go-course/watch?lesson=m04l04
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
