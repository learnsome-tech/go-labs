# Concurrent Systems Programming with Go — lesson m01l04 — Your First Program and go run
# https://learnsome.tech/courses/go-course/watch?lesson=m01l04
# © LearnSome.tech
package main

import (
	"fmt"
	"os"
)

func main() {
	unhealthy := 2
	fmt.Println("checked three targets")
	if unhealthy > 0 {
		fmt.Fprintf(os.Stderr, "%d unhealthy\n", unhealthy)
		os.Exit(1)
	}
	fmt.Println("all healthy")
}
