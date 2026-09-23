# Concurrent Systems Programming with Go — lesson m01l03 — Go Workspaces and Modules
# https://learnsome.tech/courses/go-course/watch?lesson=m01l03
# © LearnSome.tech
package main

import (
	"fmt"

	"course.local/lib"
)

func main() {
	fmt.Println(lib.Status())
}
