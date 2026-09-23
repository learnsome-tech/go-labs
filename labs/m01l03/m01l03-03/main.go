# Concurrent Systems Programming with Go — lesson m01l03 — Go Workspaces and Modules
# https://learnsome.tech/courses/go-course/watch?lesson=m01l03
# © LearnSome.tech
package main

import (
	"fmt"

	"course.local/probe/internal/check"
)

func main() {
	fmt.Println(check.Describe(check.Target{Name: "db-1", Port: 5432}))
}
