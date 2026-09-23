# Concurrent Systems Programming with Go — lesson m04l01 — Goroutines and Concurrent Execution
# https://learnsome.tech/courses/go-course/watch?lesson=m04l01
# © LearnSome.tech
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	for _, name := range []string{"api", "worker"} {
		wait.Add(1)
		go func(target string) {
			defer wait.Done()
			fmt.Println(target)
		}(name)
	}
	wait.Wait()
}
