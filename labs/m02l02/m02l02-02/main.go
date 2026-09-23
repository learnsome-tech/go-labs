# Concurrent Systems Programming with Go — lesson m02l02 — Arrays, Slices, and Maps
# https://learnsome.tech/courses/go-course/watch?lesson=m02l02
# © LearnSome.tech
package main

import "fmt"

func main() {
	items := []string{"api", "worker"}
	items = append(items, "scheduler")
	fmt.Println(len(items))
	fmt.Println(items[1])
	fmt.Println(items[:2])
}
