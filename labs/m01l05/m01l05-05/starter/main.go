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
