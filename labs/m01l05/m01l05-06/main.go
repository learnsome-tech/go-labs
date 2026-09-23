# Concurrent Systems Programming with Go — lesson m01l05 — Variables, Typing, and Zero Values
# https://learnsome.tech/courses/go-course/watch?lesson=m01l05
# © LearnSome.tech
package main

import "fmt"

func main() {
	var total int = 200
	var used int = 164
	var factor float64 = 1.25

	fmt.Println(float64(used) / float64(total) * factor)
	fmt.Println(used * int(factor))
	fmt.Println(float64(total) * factor)
}
