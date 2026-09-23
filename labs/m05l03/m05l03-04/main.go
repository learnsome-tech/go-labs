# Concurrent Systems Programming with Go — lesson m05l03 — The time Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l03
# © LearnSome.tech
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2025, time.January, 2, 3, 4, 5, 0, time.UTC)
	fmt.Println(t.Format(time.RFC3339))
}
