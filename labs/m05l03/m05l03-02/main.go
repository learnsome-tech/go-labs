# Concurrent Systems Programming with Go — lesson m05l03 — The time Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l03
# © LearnSome.tech
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Millisecond)
	fmt.Println(time.Since(start) > 0)
}
