# Concurrent Systems Programming with Go — lesson m04l04 — The context Package
# https://learnsome.tech/courses/go-course/watch?lesson=m04l04
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
)

func report(ctx context.Context) {
	if value := ctx.Value("request"); value != nil {
		fmt.Println(value)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "request", "probe")
	report(ctx)
}
