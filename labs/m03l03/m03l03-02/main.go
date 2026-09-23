# Concurrent Systems Programming with Go — lesson m03l03 — Panics and Recover
# https://learnsome.tech/courses/go-course/watch?lesson=m03l03
# © LearnSome.tech
package main

import "fmt"

func lookup(items []string, index int) (string, error) {
	if index < 0 || index >= len(items) {
		return "", fmt.Errorf("index outside list")
	}
	return items[index], nil
}

func main() {
	value, err := lookup([]string{"api"}, 4)
	if err != nil {
		fmt.Println("lookup failed")
		return
	}
	fmt.Println(value)
}
