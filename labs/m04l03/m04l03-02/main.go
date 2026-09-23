# Concurrent Systems Programming with Go — lesson m04l03 — The select Statement
# https://learnsome.tech/courses/go-course/watch?lesson=m04l03
# © LearnSome.tech
package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)
	go func() { time.Sleep(time.Millisecond); result <- "ready" }()
	select {
	case value := <-result:
		fmt.Println(value)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
