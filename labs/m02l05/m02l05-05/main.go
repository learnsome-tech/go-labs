# Concurrent Systems Programming with Go — lesson m02l05 — Type Assertions and the Empty Interface
# https://learnsome.tech/courses/go-course/watch?lesson=m02l05
# © LearnSome.tech
package main

import "fmt"

func timeout(value any) (int, bool) {
	seconds, ok := value.(int)
	if !ok || seconds < 1 {
		return 0, false
	}
	return seconds, true
}

func main() {
	if seconds, ok := timeout(5); ok {
		fmt.Println("timeout", seconds)
	}
	if _, ok := timeout("five"); !ok {
		fmt.Println("invalid timeout")
	}
}
