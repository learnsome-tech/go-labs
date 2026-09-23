# Concurrent Systems Programming with Go — lesson m06l04 — Building Static Binaries
# https://learnsome.tech/courses/go-course/watch?lesson=m06l04
# © LearnSome.tech
package main

import (
	"fmt"
	"runtime
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
