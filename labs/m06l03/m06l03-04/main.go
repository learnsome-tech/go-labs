# Concurrent Systems Programming with Go — lesson m06l03 — Graceful Shutdown on SIGTERM
# https://learnsome.tech/courses/go-course/watch?lesson=m06l03
# © LearnSome.tech
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	fmt.Println(err)
}
