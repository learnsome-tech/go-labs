# Concurrent Systems Programming with Go — lesson m07l02 — Executing Concurrent Checks
# https://learnsome.tech/courses/go-course/watch?lesson=m07l02
# © LearnSome.tech
package main

import (
	"fmt"
	"sync"
)

func main() {
	results := make(chan string, 2)
	var wait sync.WaitGroup
	for _, name := range []string{"api", "web"} {
		wait.Add(1)
		go func(target string) {
			defer wait.Done()
			results <- target + " checked"
		}(name)
	}
	wait.Wait(); close(results)
	for result := range results { fmt.Println(result) }
}
