# Concurrent Systems Programming with Go — lesson m01l05 — Variables, Typing, and Zero Values
# https://learnsome.tech/courses/go-course/watch?lesson=m01l05
# © LearnSome.tech
package main

import "fmt"

var release = "stable"

func main() {
	host := "db-1"
	port := 5432
	ratio := 0.75
	var retries int64 = 3
	var enabled bool

	fmt.Printf("%T %T %T\n", host, port, ratio)
	fmt.Printf("%T %T %T\n", retries, enabled, release)
	fmt.Println(host, port, ratio, retries, enabled, release)
}
