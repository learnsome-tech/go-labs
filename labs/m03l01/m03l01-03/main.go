# Concurrent Systems Programming with Go — lesson m03l01 — Errors as Values
# https://learnsome.tech/courses/go-course/watch?lesson=m03l01
# © LearnSome.tech
package main

import "fmt"

func load(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is empty")
	}
	return "loaded " + name, nil
}

func main() {
	value, err := load("")
	if err != nil {
		fmt.Println("load failed:", err)
		return
	}
	fmt.Println(value)
}
