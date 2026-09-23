# Concurrent Systems Programming with Go — lesson m07l03 — Handling Timeouts and Errors
# https://learnsome.tech/courses/go-course/watch?lesson=m07l03
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := ctx.Err(); err != nil {
		fmt.Println("unhealthy:", err)
	}
}
