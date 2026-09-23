# Concurrent Systems Programming with Go — lesson m01l05 — Variables, Typing, and Zero Values
# https://learnsome.tech/courses/go-course/watch?lesson=m01l05
# © LearnSome.tech
package main

import "fmt"

type state int

const (
	unknown state = iota
	healthy
	degraded
	down
)

const maxProbes = 3

func main() {
	current := degraded
	fmt.Println(unknown, healthy, degraded, down)
	fmt.Println(current == degraded, current > healthy)
	fmt.Printf("%T %d of %d\n", current, current, maxProbes)
}
