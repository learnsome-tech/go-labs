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
	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("worker finished")
	}()
	wait.Wait()
	fmt.Println("main finished")
}
