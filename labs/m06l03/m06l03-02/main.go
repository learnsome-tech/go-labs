# Concurrent Systems Programming with Go — lesson m06l03 — Graceful Shutdown on SIGTERM
# https://learnsome.tech/courses/go-course/watch?lesson=m06l03
# © LearnSome.tech
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	fmt.Println("ready")
	<-stop
	fmt.Println("stopping")
}
